package concurrency

ch := make(chan int)

go func() {
    ch <- 42
}()

val := <-ch
fmt.Println(val)
