package concurrency

import (
	"fmt"
)

func UnbufferedExample() {
	ch := make(chan int)

	go func() {
		ch <- 42 // blocks until receiver is ready
	}()

	v := <-ch
	fmt.Println("received:", v)
}
