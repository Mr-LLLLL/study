package main

import (
    "fmt"
)

func main() {
    var x, y int
    fmt.Println("input two number:")
    fmt.Scanf("%d,%d", &x, &y)
    fmt.Printf("%d and %d greatest common divisor: %d\n", x, y, gcd(x, y))
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}
