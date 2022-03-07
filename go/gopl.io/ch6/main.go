package main

import (
	"fmt"
	"net/url"

	"gopl.io/ch2/tempconv"
	"gopl.io/ch6/geometry"
	"gopl.io/ch6/intset"
)

func main() {
	perim := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	p1 := geometry.Point{X: 1, Y: 2}
	fmt.Printf("%T\n", p1)

	fmt.Println(perim.Distance())

	s := []int{1, 2, 3, 4}
	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println(perim)
	fmt.Printf("%T\n", perim)

	fmt.Println(geometry.Point{X: 1, Y: 2}.Distance(geometry.Point{X: 2, Y: 3}))
	p := geometry.Point{X: 1, Y: 2}
	p.ScaleBy(2)
	// 指针类型方法能用
	// geometry.Point{1, 2}.ScaleBy(2)
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	// m.Add("item", "3")

	var m2 map[int]int
	m1 := make(map[int]int)
	m1[1] = 1
	m1[2] = 2
	m2 = m1
	fmt.Println(m2)

	fmt.Println()
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(&x)

	y.Add(9)
	y.Add(42)
	fmt.Printf("%T, %[1]s\n", y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))

	c := tempconv.Celsius(1)
	fmt.Println(c.String())

	var i = [64]int{1, 2}
	s1 := i[:0]
	var s2 []int
	if s1 == nil {
		fmt.Println("s1 == nil")
	}
	if s2 == nil {
		fmt.Println("s2 == nil")
	}
	fmt.Println(s1, s2)

	type test int
	t := test(1)
	fmt.Printf("%T\n", t)
}
