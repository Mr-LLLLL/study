package main

import "fmt"

type queen struct {
    x, y int
}

func main() {
    var n int
    for true {
        fmt.Print("please input width of cheese:")
        fmt.Scan(&n)
        fmt.Println("the possible cheese is:")
        placeQueen(n)
    }
}

func placeQueen(n int) {
    var solu []queen
    q := queen{0, 0}
    for q.x > 0 || q.y < n {
        if len(solu) >= n || q.y >= n {
            q = solu[len(solu) - 1]
            q.y++
            solu = solu[:len(solu) - 1]
        } else {
            for q.y < n && isExist(solu, q) {
                q.y++
            }
            if (q.y < n) {
                solu = append(solu, q)
                if len(solu) >= n {
                    for _, v := range(solu) {
                        fmt.Print(v)
                    }
                    fmt.Println()
                }
                q.x++
                q.y = 0
            }
        }
    }

}

func equal(q, q1 queen) bool {
    return q.x == q1.x ||
            q.y == q1.y ||
            q.x + q.y == q1.x + q1.y ||
            q.x - q.y == q1.x - q1.y
}

func isExist(stack []queen, q queen) bool {
    for _, v := range(stack) {
        if equal(v, q) {
            return true
        }
    }
    return false
}
