package wall

import "maze/base"

type IWall interface {
	base.MapSite
	Clone() IWall
}

type NormalWall struct{}

func (w *NormalWall) Enter() {
	panic("not implemented") // TODO: Implement
}

func (w *NormalWall) Clone() IWall {
	return new(NormalWall)
}

func NewNormalWall() *NormalWall {
	return new(NormalWall)
}

type BombedWall struct {
	_bomb bool
}

func (b *BombedWall) Enter() {
	panic("not implemented")
}

func (b *BombedWall) Clone() IWall {
	return &BombedWall{
		_bomb: b._bomb,
	}
}

func (b *BombedWall) HasBomb() bool {
	return b._bomb
}

func NewBombedWall() *BombedWall {
	return &BombedWall{
		_bomb: true,
	}
}
