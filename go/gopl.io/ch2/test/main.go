package main

import "fmt"

func Bar() string {
	return "bar"
}
func Foo() string {
	return "foo"
}

func Qux(v string) string {
	if v == "foo" {
		return Foo()
	}

	if v == "bar" {
		return Bar()
	}

	return "INVALID"
}

func main() {
	var b int
	var ab int
	foo_hello := 1
	fmt.Println("Hello GopherCon")
	fmt.Println("Hello GopherCon")
	fmt.Println(foo_hello)
	fmt.Println(b)
	fmt.Println(ab)
	i := 1 + 2
	fmt.Println(i)

	var f float32 = 16777216
	fmt.Println(f == f+1)
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}
