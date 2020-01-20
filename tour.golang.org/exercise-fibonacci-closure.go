package main

import "fmt"

func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        next := a
        a, b = b, a+b
        return next
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
