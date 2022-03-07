package main

import (
	"di/di"
	"fmt"
)

func main() {
	container := di.New()
	if err := container.Provide(NewA); nil != err {
		panic(err)
	}
	if err := container.Provide(NewB); nil != err {
		panic(err)
	}
	if err := container.Provide(NewC); nil != err {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.B.C.Num)
	})
	if nil != err {
		panic(err)
	}
}

type A struct {
	B *B
}

func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

type B struct {
	C *C
}

func NewB(c *C) *B {
	return &B{
		C: c,
	}
}

type C struct {
	Num int
}

func NewC() *C {
	return &C{
		Num: 1,
	}
}
