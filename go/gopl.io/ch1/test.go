package main

import (
    "fmt"


func main() {
    fmt.Println(signum(5))
}

func signum(x int) int {
    switch {
    case x > 0:
        return +1
    default:
        return 0
    case x < 0:
        return -1
    }
}
