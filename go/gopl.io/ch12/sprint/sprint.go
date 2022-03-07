package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func main() {
	var m my_class
	fmt.Println(Sprint(m))

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	t := reflect.TypeOf(3)
	fmt.Printf("%T\n", t)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	fmt.Printf("%T\n", v)

	t1 := v.Type()
	fmt.Println(t1.String())
	fmt.Printf("%T\n", t1)
}

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	type my_funcer interface {
		my_func() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case my_funcer:
		return x.my_func()
	case string:
		fmt.Println("string")
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

type my_class struct {
	i int
	s string
}

func (_ my_class) String() string {
	return "hello"
}

func (_ my_class) my_func() string {
	return "world"
}
