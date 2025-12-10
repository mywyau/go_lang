package main

import (
    "fmt"
    "go_lang/mini_concurrency_app/internal/worker"
    "time"
)

func main() {
    wp := worker.NewWorkerPool(3)
    wp.Start()

    // Submit some jobs
    go func() {
        for i := 1; i <= 10; i++ {
            wp.Submit(worker.Job(i))
        }
        wp.Stop()
    }()

    // Read results until channel closes
    for res := range wp.Results() {
        fmt.Println(res)
    }

    // Give workers time to shut down cleanly (optional)
    time.Sleep(100 * time.Millisecond)
    fmt.Println("All done!")
}
