package concurrency

import (
	"context"
	"fmt"
	"time"
)

func ContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker stopped")
				return
			default:
				ch <- "working..."
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	cancel() // signal cancellation
	time.Sleep(50 * time.Millisecond)
}
