package concurrency

import (
	"fmt"
	"time"
)

func SelectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "first"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "second"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("from ch2:", msg)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("timeout")
	}
}
