package adapter4object

import (
	"adapter/base"
	"adapter/shape"
)

type TextShape struct {
	shape.Shape
	_text *shape.TextView
}

func (t *TextShape) BoundingBox(bottomLeft, topRight *base.Point) {
	var bottom, left, width, height *base.Coord

	t._text.GetOrigin(bottom, left)
	t._text.GetExtent(width, height)

	t.BottomLeft = base.NewPoint(*bottom, *left)
	t.TopRight = base.NewPoint(*bottom+*height, *left+*width)
}

func (TextShape) CreateManipulator() *base.Manipulator {
	return new(base.Manipulator)
}

func (t *TextShape) IsEmpty() bool {
	return t._text.IsEmpty()
}

func NewTextShape(t *shape.TextView) TextShape {
	return TextShape{
		_text: t,
	}
}
