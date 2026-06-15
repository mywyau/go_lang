package concurrency

import (
	"fmt"
)

/*
Unbuffered channel example

Idea:
- An unbuffered channel has no storage space.
- A send blocks until another goroutine is ready to receive.
- A receive blocks until another goroutine is ready to send.
- This makes unbuffered channels useful for synchronising goroutines.

In this example:
1. main creates an unbuffered int channel.
2. a goroutine tries to send the value 42 into the channel.
3. that send waits until the main goroutine receives from the channel.
4. main receives the value and prints it.

Important:
- If we tried to send on the channel without a separate goroutine,
  the program would deadlock because there would be no receiver ready.
*/


func UnbufferedExample() {
	ch := make(chan int)

	go func() {
		ch <- 42 // blocks until receiver is ready
	}()

	v := <-ch
	fmt.Println("received:", v)
}