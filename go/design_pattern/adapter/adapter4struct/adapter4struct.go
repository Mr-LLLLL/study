package adapter4struct

import (
	"adapter/base"
	"adapter/shape"
)

type TextShape struct {
	shape.Shape
	shape.TextView
}

func (t *TextShape) BoundingBox(bottomLeft, topRight *base.Point) {
	var bottom, left, width, height *base.Coord

	t.GetOrigin(bottom, left)
	t.GetExtent(width, height)

	t.BottomLeft = base.NewPoint(*bottom, *left)
	t.TopRight = base.NewPoint(*bottom+*height, *left+*width)
}

func (TextShape) CreateManipulator() *base.Manipulator {
	return new(base.Manipulator)
}

func (TextShape) GetOrigin(x, y *base.Coord) {}

func (TextShape) GetExtent(width, height *base.Coord) {}

func (t *TextShape) IsEmpty() bool {
	return t.TextView.IsEmpty()
}
