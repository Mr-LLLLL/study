package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"gopl.io/ch7/bytecounter"
	"gopl.io/ch7/tempconv"
)

func main() {
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	// var w io.Writer
	// w = os.Stdout
	// w = new(bytes.Buffer)
	os.Stdout.Write([]byte("hello\n"))

	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	// any = new(bytes.Buffer)
	fmt.Println(any)

	var s1 string
	var s string = "hello"
	fmt.Sscanf(s, "%s", &s1)
	fmt.Println(s, s1)

	flag.Parse()
	fmt.Println(*temp)

	var x interface{} = 1
	var x1 interface{} = 2
	fmt.Printf("%T\n", x)
	fmt.Println(x == x1)

	fmt.Println("==================================================")
	fmt.Println(test1())

	var w io.Writer
	w = os.Stdout
	w = w.(io.Writer)
	fmt.Printf("%T\n", w)

	var c1 inter
	var c2 class = class{[]string{"hello"}}
	c1 = &c2
	c3 := c1.(*class)
	fmt.Println(c1, c2)
	fmt.Printf("%T %p %p\n", c3, c3, &c2)

	var s3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d, cap=%d, %v\n", len(s3), cap(s3), s3[1:3])
	var s4 = []string{"1", "2", "3"}
	fmt.Println(s4)
	fmt.Println(strings.Join(s4, " , "))
}

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

type inter interface {
	testfunc() int
}

type class struct {
	s []string
}

func (c class) testfunc() int {
	return 2
}

func test1() (_ int) {
	return 1
}
