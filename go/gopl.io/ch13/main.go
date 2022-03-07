package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 1
	var in interface{} = i
	fmt.Println(unsafe.Sizeof(in))
	fmt.Println(unsafe.Sizeof(""))

	fmt.Println(unsafe.Sizeof(struct {
		bool
		float64
		int16
	}{}), unsafe.Sizeof(struct {
		float64
		int16
		bool
	}{}))
	fmt.Println()
	fmt.Println(unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Println(unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Println(unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))

	fmt.Println()
	fmt.Printf("%#016x\n", Float64bits(0.0))
	fmt.Printf("%#016x\n", uint64(float64(1.0)))

	fmt.Println()
	var ss []string
	var si []int
	ss1 := []string{}
	fmt.Println(unsafe.Sizeof(ss), unsafe.Sizeof(si), unsafe.Sizeof(ss1))

	fmt.Println()
	m1 := map[int]string{1: "hello", 2: "world"}
	for i := 0; i < 1000; i++ {
		m1[i] = "hello"
	}
	m2 := m1
	fmt.Printf("%+v", m2)

	fmt.Println()
	s1 := []int{1, 2, 3}
	s2 := s1
	for i := 0; i < 1000; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("%v", s2)

	type MyType int
	fmt.Println()
	fmt.Printf("%T\n", MyType(1))
}

var x struct {
	a bool
	b int16
	c []int
}

func Float64bits(f float32) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
