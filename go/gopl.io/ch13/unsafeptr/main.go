package main

import (
	"fmt"
	"unsafe"
)

func main() {
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)

	// sometime wrong, because the garbage recollect, uintptr is a value, when gc start, uintptr don't change(other pointer point a variable will be repoint a new address)
	// fmt.Println()
	// tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	// pb1 := (*int16)(unsafe.Pointer(tmp))
	// *pb1 = 43
	// fmt.Println(x.b)
}

var x struct {
	a bool
	b int16
	c []int
}
