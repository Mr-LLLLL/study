package main

import (
	"fmt"
	"log"
)

func main() {
    var s string
    for true {
        fmt.Printf("please input string:")
        _, err := fmt.Scan(&s)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("the string is paren:%t\n", isParen(s))
    }
}

func isParen(s string) bool {
    var stack []byte
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case '(', '[', '{':
            stack = append(stack, s[i])
        case ')':
            if len(stack) == 0 || stack[len(stack) - 1] != '(' {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
        case ']':
            if len(stack) == 0 || stack[len(stack) - 1] != '[' {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
        case '}':
            if len(stack) == 0 || stack[len(stack) - 1] != '{' {
                return false
            } else {
                stack = stack[:len(stack) - 1]
            }
        }
    }
    return len(stack) == 0
}
