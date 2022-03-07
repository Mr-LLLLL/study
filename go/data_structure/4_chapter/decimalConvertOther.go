package main

import (
	"container/list"
	"fmt"
	"log"
)

func main() {
	var n, b int
	var s list.List
	for true {
		fmt.Println("please input number and base:")
		_, err := fmt.Scan(&n, &b)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Printf("the number of %d base is: 0x", b)
		convert(&s, n, b)
		for s.Len() != 0 {
			fmt.Printf("%c", s.Back().Value)
			s.Remove(s.Back())
		}
		fmt.Println()
	}
}

var digit = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func convert(s *list.List, n, b int) {
	for n > 0 {
		s.PushBack(digit[n%b])
		n /= b
	}
}
