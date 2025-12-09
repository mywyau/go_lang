package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func MutexExample() {
	var mu sync.Mutex
	count := 0

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println("count:", count)
}
