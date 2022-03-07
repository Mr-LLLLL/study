package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(links.Extract("http://www.baidu.com"))
	fmt.Printf("%*v\n", 5, []int{1, 2, 3, 4})

	// in := bufio.NewReader(os.Stdin)
	// for {
	//     r, err := in.ReadByte()
	//     if err == io.EOF {
	//         break
	//     }
	//     if err != nil {
	//         fmt.Printf("read failed: %v", err)
	//     }
	//     println(r)
	// }
	add1 := func(r rune) rune {
		return r + 1
	}

	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))

	var s []string
	s = append(s, "hello")
	s1 := s
	s1[0] = "nihao"
	s1 = append(s, "world", "yes", "no")
	s1[0] = "hello"
	fmt.Println(s)
	fmt.Println(s1)
	err := os.MkdirAll("/hello/world", 0755)
	if err != nil {
		fmt.Println(err)
	}
	// go closure capture a variable address not a value
	i := 1
	f := func() int {
		return i + 1
	}
	i++
	fmt.Println(f())
	values := []int{1, 2, 3, 4}
	fmt.Println(values)
	_ = sum(values...)

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s%s", name, "hello")
	testPanic()
	fmt.Println("not end")
}

func testPanic() {
	fmt.Println("test...........")
	i := 1
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("nothing a")
		case "hello":
			i++
			fmt.Println("world", i)
		default:
			fmt.Println("i don know")
		}
	}()
	if true {
		panic("hello")
	}
	fmt.Println("func you")
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
	fmt.Printf("%T%T\n", args...)
}
