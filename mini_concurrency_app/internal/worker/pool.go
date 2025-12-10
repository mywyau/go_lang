package worker

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Job int

type WorkerPool struct {
    workers int
    jobs    chan Job
    results chan string
    ctx     context.Context
    cancel  context.CancelFunc
    wg      sync.WaitGroup // NEW: tracks worker completion
}

func NewWorkerPool(workerCount int) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())

    wp := &WorkerPool{
        workers: workerCount,
        jobs:    make(chan Job),
        results: make(chan string),
        ctx:     ctx,
        cancel:  cancel,
    }

    return wp
}

func (wp *WorkerPool) Start() {
    for w := 1; w <= wp.workers; w++ {
        wp.wg.Add(1)
        go wp.worker(w)
    }

    // Once all workers finish, close results
    go func() {
        wp.wg.Wait()
        close(wp.results)
    }()
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()

    for {
        select {
        case <-wp.ctx.Done():
            return
        case job, ok := <-wp.jobs:
            if !ok {
                return // jobs channel closed → worker stops
            }
            // simulate work
            time.Sleep(150 * time.Millisecond)
            wp.results <- fmt.Sprintf("worker %d processed job %d", id, job)
        }
    }
}

func (wp *WorkerPool) Submit(job Job) {
    wp.jobs <- job
}

func (wp *WorkerPool) Results() <-chan string {
    return wp.results
}

func (wp *WorkerPool) Stop() {
    close(wp.jobs) // workers exit when jobs channel is closed
    wp.cancel()    // context stops active workers nicely
}
