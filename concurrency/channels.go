package concurrency

import "fmt"

func BasicChannel() {
	ch := make(chan string, 3) // buffer of 3

	go func() {
		ch <- "hello"
		ch <- "from"
		ch <- "goroutine"
		close(ch) // sender closes the channel
	}()

	for msg := range ch { // drains until closed
		fmt.Println(msg)
	}
}
