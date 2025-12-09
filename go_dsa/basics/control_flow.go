package basics

import (
    "fmt"
    "time"
)

func IfExample(x int) {
    if x > 10 {
        fmt.Println("big")
    } else {
        fmt.Println("small")
    }
}

func IfInitExample() {
    if y := time.Now().Hour(); y < 12 {
        fmt.Println("morning")
    } else {
        fmt.Println("afternoon or evening")
    }
}