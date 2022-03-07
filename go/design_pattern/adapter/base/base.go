package base

type Point struct {
	x, y int
}

type Manipulator struct{}

type Coord int

func NewPoint(x, y Coord) Point {
	return Point{
		x: int(x),
		y: int(y),
	}
}
