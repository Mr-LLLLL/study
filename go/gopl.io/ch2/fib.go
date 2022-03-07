package main

import "fmt"

func main() {
    var n int
    fmt.Println("input which number you want to get fibonacci:")
    fmt.Scanln(&n)
    fmt.Printf("%dth fibonacci is: %d\n", n, fib(n))
}

func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x + y
    }
    return x
}
