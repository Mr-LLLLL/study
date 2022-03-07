package strategy

import (
	"fmt"
	"strategy/base"
)

type Composition struct {
	_componentCount int
	_lineWidth      int
	_lineBreaks     *int
	_lineCount      int

	_compositor ICompositor
	_components *base.Component
}

func (c Composition) Repair() {
	var (
		natural        []base.Coord
		stretchability []base.Coord
		shrinkability  []base.Coord
		componentCount int
		breaks         []int
		breakCount     int
	)

	breakCount = c._compositor.Compose(natural, stretchability, shrinkability, componentCount, c._lineWidth, breaks)
	fmt.Println(breakCount)
}

func NewComposition(c ICompositor) *Composition {
	return &Composition{
		_compositor: c,
	}
}

type ICompositor interface {
	Compose(natural, stretch, shrink []base.Coord, componentCount, linewidth int, breaks []int) int
}

type SimpleCompositor struct{}

func (SimpleCompositor) Compose(natural, stretch, shrink []base.Coord, componentCount, linewidth int, breaks []int) int {
	return 0
}

type TeXCompositor struct{}

func (TeXCompositor) Compose(natural, stretch, shrink []base.Coord, componentCount, linewidth int, breaks []int) int {
	return 0
}

type ArrayCompositor struct{}

func (ArrayCompositor) Compose(natural, stretch, shrink []base.Coord, componentCount, linewidth int, breaks []int) int {
	return 0
}
