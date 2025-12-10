package processor

import (
    "context"
    "io"
    "os"
    "sync"
)

type Job struct {
    Path string
}

type Result struct {
    Path string
    Size int64
    Err  error
}

type Pool struct {
    workers int
    jobs    chan Job
    results chan Result
    ctx     context.Context
    cancel  context.CancelFunc
    wg      sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
    ctx, cancel := context.WithCancel(context.Background())
    return &Pool{
        workers: workerCount,
        jobs:    make(chan Job),
        results: make(chan Result),
        ctx:     ctx,
        cancel:  cancel,
    }
}

func (p *Pool) Start() {
    for w := 1; w <= p.workers; w++ {
        p.wg.Add(1)
        go p.worker(w)
    }

    // When workers finish, close results
    go func() {
        p.wg.Wait()
        close(p.results)
    }()
}

func (p *Pool) worker(id int) {
    defer p.wg.Done()

    for {
        select {
        case <-p.ctx.Done():
            return

        case job, ok := <-p.jobs:
            if !ok {
                return
            }

            size, err := processFile(job.Path)
            p.results <- Result{
                Path: job.Path,
                Size: size,
                Err:  err,
            }
        }
    }
}

func (p *Pool) Submit(j Job) {
    p.jobs <- j
}

func (p *Pool) Results() <-chan Result {
    return p.results
}

func (p *Pool) Stop() {
    close(p.jobs)
    p.cancel()
}

// ---------- FILE PROCESSING LOGIC ----------

func processFile(path string) (int64, error) {
    f, err := os.Open(path)
    if err != nil {
        return 0, err
    }
    defer f.Close()

    // Count bytes
    n, err := io.Copy(io.Discard, f)
    return n, err
}
