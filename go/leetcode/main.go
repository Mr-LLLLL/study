package main

import (
	"fmt"
	"sort"
	"time"
)

var defaultTimeFormat = "2006-01-02 15:04:05"         // 字符串格式化为时间的格式
var localTime, _ = time.LoadLocation("Asia/Shanghai") // 上海时区

func main() {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
		ch <- 2
		fmt.Println("hello")
	}()

	go func() {
		fmt.Println(<-ch)
		ch <- 3
		fmt.Println(<-ch)
	}()

	select {}
}

func f3(int) {

}

func f2(arr []int) {
	arr[0], arr[1] = arr[1], arr[0]
}

func f1() {
	ia := [...]int{1, 2, 3, 4, 5}
	ia2 := ia[1:3]
	ia2 = append(ia2, 6, 7, 8, 9)
	fmt.Println(ia, ia2)
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 1073741824.\n")
	answer := sort.Search(1073741824, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
