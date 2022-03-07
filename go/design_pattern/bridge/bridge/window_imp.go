package bridge

import "bridge/base"

type IWindowImp interface {
	DeviceRect(c1, c2, c3, c4 base.Coord)
}

type XWindowImp struct{}

func (w XWindowImp) DeviceRect(c1, c2, c3, c4 base.Coord) {
	// TODO implement
}

type PMWindowImp struct{}

func (w PMWindowImp) DeviceRect(c1, c2, c3, c4 base.Coord) {
	// TODO implement
}
