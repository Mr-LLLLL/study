package main

import (
	"fmt"
	"image/color"

	"gopl.io/ch6/geometry"
)

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{geometry.Point{X: 1, Y: 1}, red}
	var q = ColoredPoint{geometry.Point{X: 5, Y: 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	// r := new(Rocket)
	// time.AfterFunc(1*time.Second, r.Launch)
	// time.Sleep(2 * time.Second)

	p1 := geometry.Point{X: 1, Y: 2}
	q1 := geometry.Point{X: 4, Y: 6}

	distance := geometry.Point.Distance
	fmt.Println(distance(p1, q1))
	fmt.Printf("%T\n", distance)

	scale := (*geometry.Point).ScaleBy
	scale(&p1, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

type Rocket struct{}

type R = Rocket

func (r *R) Launch() {
	fmt.Println("launch.....")
}
