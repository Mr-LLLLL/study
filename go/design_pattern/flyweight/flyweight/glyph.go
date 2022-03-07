package flyweight

import "flyweight/base"

type IGlyph interface {
	Draw(*base.Window, *GlyphContext)
	SetFont(*base.Font, *GlyphContext)
	GetFont() *base.Font

	First(*GlyphContext)
	Next(*GlyphContext)
	IsDone(*GlyphContext) bool
	Current(*GlyphContext) IGlyph

	Insert(IGlyph, *GlyphContext)
	Remove(*GlyphContext)
}

type Glyph struct{}

func (g *Glyph) Draw(_ *base.Window, _ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) SetFont(_ *base.Font, _ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) GetFont() *base.Font {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) First(_ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) Next(_ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) IsDone(_ *GlyphContext) bool {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) Current(_ *GlyphContext) IGlyph {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) Insert(_ IGlyph, _ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Glyph) Remove(_ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

type Charater struct {
	_charcode byte
}

func (g *Charater) Draw(_ *base.Window, _ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) SetFont(_ *base.Font, _ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) GetFont() *base.Font {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) First(_ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) Next(_ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) IsDone(_ *GlyphContext) bool {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) Current(_ *GlyphContext) IGlyph {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) Insert(_ IGlyph, _ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

func (g *Charater) Remove(_ *GlyphContext) {
	panic("not implemented") // TODO: Implement
}

type GlyphContext struct {
	_index int
	_fonts *base.BTree
}

func (g *GlyphContext) Next(step int) {}

func (g *GlyphContext) Insert(quantity int) {}

func (g *GlyphContext) GetFont() *base.Font {
	panic("not implemented")
}

func (g *GlyphContext) SetFont(_ *base.Font, spen int) {}

type GlyphFactory struct {
	_character [128]*Charater
}

func (g *GlyphFactory) CreateCharater(c byte) *Charater {
	if g._character[c]._charcode == 0 {
		g._character[c]._charcode = c
	}

	return g._character[c]
}

func (g *GlyphFactory) CreateRow() *base.Row {
	return new(base.Row)
}

func (g *GlyphFactory) CreateColumn() *base.Column {
	return new(base.Column)
}
