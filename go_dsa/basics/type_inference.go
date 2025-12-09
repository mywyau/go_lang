package basics

import "fmt"

func TypeInferenceExample(nums []int, target int) int {

    x := 10          // inferred as int
    name := "John"   // inferred as string
    alive := true    // inferred as bool

    fmt.Println(x, name, alive)

    // return something just to satisfy the function signature
    return x + target
}
