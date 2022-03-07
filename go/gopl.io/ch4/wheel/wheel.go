package wheel

import (
	"fmt"
	"sort"
	"time"
)

func test() {
	var a [3]int = [...]int{1, 3, 4}
	fmt.Println(a)
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

	q := [...]int{1, 2, 3}
	q = [3]int{1, 2}
	fmt.Printf("%T\n", q)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "(", GBP: "@", RMB: "!"}
	fmt.Println(RMB, symbol[RMB])

	var s []int
	fmt.Printf(" slice test len = %d, cap = %d\n", len(s), cap(s))
	fmt.Println(s == nil)
	s = nil
	fmt.Println(s == nil)
	s = []int(nil)
	fmt.Println(s == nil)
	s = []int{}
	fmt.Println(s == nil)
	s = make([]int, 0)
	fmt.Println(s == nil)

	s1 := [100]int{1, 2, 3}
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)

	s = make([]int, 10)
	fmt.Println(s)

	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
		fmt.Printf("len = %d, cap = %d, addr = %x\n", len(runes), cap(runes), runes)
	}
	fmt.Printf("%q\n", runes)
	runes1 := []rune{1, 2, 3, 4, 5, 6}

	runes1 = runes1[:3]
	fmt.Printf("%q, len = %d, cap = %d\n", runes1, len(runes1), cap(runes1))
	runes2 := runes1[1:5]
	fmt.Printf("%q, len = %d, cap = %d\n", runes2, len(runes2), cap(runes2))

	copy(runes1, runes)
	runes1[1] = 'y'
	fmt.Printf("=====%q, %q\n", runes, runes1)

	var var1 int
	var var2 int
	var var3 int
	var4 := 0
	var5 := 0
	var6 := 0
	fmt.Printf("var1 addr = %x\n", &var1)
	fmt.Printf("var2 addr = %x\n", &var2)
	fmt.Printf("var3 addr = %x\n", &var3)
	fmt.Printf("var4 addr = %x\n", &var4)
	fmt.Printf("var5 addr = %x\n", &var5)
	fmt.Printf("var6 addr = %x\n", &var6)

	var ages map[string]int = make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	delete(ages, "alice1")
	fmt.Println(ages, ages1)

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var ages2 map[string]int
	fmt.Printf("len = %d", len(ages2))

	if age, ok := ages["bob"]; !ok {
		fmt.Printf("bob is not exist!, age = %d\n", age)
	}

	var dilbert Employee
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert)

	var w Wheel = Wheel{Circle{point{8, 8}, 5, 1}, 20}
	fmt.Println(w)
	w = Wheel{
		Circle: Circle{
			point:  point{x: 8, Y: 8},
			Radius: 5,
			X:      3,
		},
		Spokes: 20,
	}
	fmt.Println(w)
	fmt.Printf("%#v\n", w)
	w.x = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)
	fmt.Printf("%#v\n", w)
}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salery    int
	ManagerID int
}

type point struct {
	x, Y int
}

type Circle struct {
	point
	Radius int
	X      int
}

type Wheel struct {
	Circle
	Spokes int
}
