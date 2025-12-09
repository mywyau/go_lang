package concurrency

import (
    "context"
    "fmt"
    "time"
)

// 1. Unbuffered channel example
func UnbufferedChannelExample() {
    ch := make(chan int)

    go func() {
        ch <- 42
    }()

    val := <-ch
    fmt.Println("Unbuffered received:", val)
}

// 2. Basic buffered channel example
func BasicChannelDemo() {
    ch := make(chan string, 3)

    go func() {
        ch <- "hello"
        ch <- "from"
        ch <- "goroutine"
        close(ch)
    }()

    for msg := range ch {
        fmt.Println("Buffered:", msg)
    }
}

// 3. Select example
func SelectChannelExample() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "message from ch1"
    }()

    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "message from ch2"
    }()

    select {
    case msg := <-ch1:
        fmt.Println("Select got:", msg)
    case msg := <-ch2:
        fmt.Println("Select got:", msg)
    case <-time.After(300 * time.Millisecond):
        fmt.Println("Select: timeout")
    }
}

// 4. Detecting closed channel
func ChannelCloseDetection() {
    ch := make(chan int, 2)
    ch <- 10
    ch <- 20
    close(ch)

    for {
        val, ok := <-ch
        if !ok {
            fmt.Println("Channel closed!")
            break
        }
        fmt.Println("CloseDetect got:", val)
    }
}

// 5. Fan-out / Fan-in example
func FanOutFanIn() {
    numbers := []int{1, 2, 3, 4, 5}
    out := make(chan int)

    // fan-out: launch workers
    for _, n := range numbers {
        go func(v int) {
            out <- v * v
        }(n)
    }

    // fan-in: collect all results
    for range numbers {
        fmt.Println("FanIn:", <-out)
    }
}

// 6. Context cancellation example
func ChannelWithContext() {
    ctx, cancel := context.WithCancel(context.Background())
    ch := make(chan string)

    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("Context worker stopped")
                return
            case ch <- "working...":
                time.Sleep(50 * time.Millisecond)
            }
        }
    }()

    fmt.Println("Context:", <-ch)
    fmt.Println("Context:", <-ch)

    cancel()
    time.Sleep(50 * time.Millisecond)
}

// 7. Directional channel example
func producer(out chan<- int) {
    for i := 1; i <= 3; i++ {
        out <- i
    }
    close(out)
}

func consumer(in <-chan int) {
    for v := range in {
        fmt.Println("consumed:", v)
    }
}

func DirectionalChannelsExample() {
    ch := make(chan int, 3)
    go producer(ch)
    consumer(ch)
}

// 8. Non-blocking send/receive example
func NonBlockingChannel() {
    ch := make(chan int, 1)

    select {
    case ch <- 99:
        fmt.Println("NonBlocking send!")
    default:
        fmt.Println("Send would block")
    }

    select {
    case val := <-ch:
        fmt.Println("NonBlocking recv:", val)
    default:
        fmt.Println("Receive would block")
    }
}
