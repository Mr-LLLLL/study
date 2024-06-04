package main

import (
	"fmt"
	"runtime"
)

type str struct {
	name string
	ss   []*str
}

type snippet struct{}

func main() {
	runtime.GOMAXPROCS(55)
	fmt.Println(runtime.GOMAXPROCS(0))
}

func (s *str) test() {
	s = nil
}
