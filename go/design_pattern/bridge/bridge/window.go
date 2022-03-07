package bridge

import "bridge/base"

type IWindow interface {
	GetWindowImpl() IWindowImp
	DrawRect(p1, p2 base.Point)
}

type Window struct {
	_imp IWindowImp
}

type IconWindow struct {
	_imp IWindowImp
}

func (w *IconWindow) GetWindowImpl() IWindowImp {
	if w._imp == nil {
		w._imp = XWindowImp{}
	}
	return w._imp
}

func (w *IconWindow) DrawRect(p1, p2 base.Point) {
	imp := w.GetWindowImpl()
	imp.DeviceRect(p1.X, p1.Y, p2.X, p2.Y)
}

type ApplicationWindow struct {
	_imp IWindowImp
}

func (w *ApplicationWindow) GetWindowImpl() IWindowImp {
	if w._imp == nil {
		w._imp = PMWindowImp{}
	}
	return w._imp
}

func (w *ApplicationWindow) DrawRect(p1, p2 base.Point) {
	imp := w.GetWindowImpl()
	imp.DeviceRect(p1.X, p1.Y, p2.X, p2.Y)
}
