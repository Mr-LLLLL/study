package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	fmt.Println(v.String())
	t := v.Type()
	fmt.Println(t.String())
	fmt.Printf("%T\n", t)
	fmt.Println(t)
	f1()

	m := map[int]int{
		1: 2,
		2: 3,
		3: 4,
	}
	var inter interface{}
	inter = m
	m1 := inter.(map[int]int)
	m1[1] = 3
	fmt.Printf("%p, %p\n", &m, &m1)
	fmt.Printf("%p, %p\n", m, m1)
	fmt.Println("============================================")
	for j := 1; j < 1000; j++ {
		m[j] = j + 3
	}
	fmt.Printf("%p, %p\n", &m, &m1)
	fmt.Printf("%p, %p\n", m, m1)
	fmt.Println(m[1])

	s1 := []int{1, 2, 3}
	s2 := s1
	fmt.Println(s1[1])
	for i := 1; i < 1000; i++ {
		s1 = append(s1, i)
	}
	s2[1] = 4
	fmt.Println("==================================================")
	fmt.Println(s1[1])

	fmt.Println("==================================================")
	x := 2
	d := reflect.ValueOf(&x).Elem()
	fmt.Printf("%T\n", d)
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(d)
	d.Set(reflect.ValueOf(4))
	fmt.Println(x)

	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(10)
	fmt.Println(x)
	rx.Set(reflect.ValueOf(3))
	fmt.Println(x)

	var y interface{}
	y = x
	ry := reflect.ValueOf(&y).Elem()
	ry.Set(reflect.ValueOf(8))
	fmt.Println(x)

	s3 := []int{1, 3, 4}
	rs3 := reflect.ValueOf(s3)
	fmt.Println(rs3.Type().String())
}

func f1() {

}
