package shape

import "adapter/base"

type IShape interface {
	BoundingBox(bottomLeft, topRight *base.Point)
	CreateManipulator() *base.Manipulator
}

type ITextView interface {
	GetOrigin(x, y *base.Coord)
	GetExtent(width, height *base.Coord)
	IsEmpty() bool
}

type Shape struct {
	BottomLeft, TopRight base.Point
}

func (Shape) BoundingBox(bottomLeft, topRight *base.Point) {
}

func (Shape) CreateManipulator() *base.Manipulator {
	return new(base.Manipulator)
}

type TextView struct{}

func (TextView) GetOrigin(x, y *base.Coord) {}

func (TextView) GetExtent(width, height *base.Coord) {}

func (TextView) IsEmpty() bool {
	return false
}
