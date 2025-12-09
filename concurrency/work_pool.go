package concurrency

import (
	"fmt"
	"time"
)

func WorkerPool() {
	
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    // Workers
    for w := 1; w <= 3; w++ {
        go func(id int) {
            for j := range jobs {
                fmt.Printf("worker %d processing job %d\n", id, j)
                time.Sleep(200 * time.Millisecond)
                results <- j * 2
            }
        }(w)
    }

    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Receive results
    for i := 0; i < 5; i++ {
        fmt.Println("result:", <-results)
    }
}
