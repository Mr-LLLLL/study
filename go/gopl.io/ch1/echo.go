package main

import (
    "fmt"
    "os"
    // "strings"
)

// func main() {
//     var s, sep string
//     for i := 0; i < len(os.Args); i++ {
//         s += sep + os.Args[i]
//         sep = " "
//     }
//     fmt.Println(s)
// }

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

// func main() {
//     fmt.Println(strings.Join(os.Args[1:], " "))
//     fmt.Println(os.Args[1:])
// }


