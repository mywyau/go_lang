package concurrency

import (
	"fmt"
	"time"
)

func BasicGoroutine() {
	go doWork() // runs asynchronously
	time.Sleep(200 * time.Millisecond)
}

func doWork() {
	fmt.Println("doing work in goroutine")
}
