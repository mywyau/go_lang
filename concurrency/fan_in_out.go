package concurrency

import (
	"fmt"
	"time"
)

func FanInFanOut() {
	work := []int{1, 2, 3, 4, 5}

	worker := func(n int, out chan<- int) {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		out <- n * n
	}

	out := make(chan int, len(work))

	for _, n := range work {
		go worker(n, out) // fan-out
	}

	// fan-in
	for i := 0; i < len(work); i++ {
		fmt.Println(<-out)
	}
}
