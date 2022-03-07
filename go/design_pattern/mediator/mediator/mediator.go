package mediator

import "reflect"

type IDialogDirector interface {
	Showdialog()
	WidgetChanged(IWidget)
	CreateWidgets()
}

type FontDialogDirector struct {
	_ok       *Button
	_cancel   *Button
	_fontList *ListBox
	_fontName *EntryField
}

func (d *FontDialogDirector) CreateWidgets() {
	d._ok = NewButton(d)
	d._cancel = NewButton(d)
	d._fontList = NewListBox(d)
	d._fontName = NewEntryField(d)
}

func (d *FontDialogDirector) WidgetChanged(w IWidget) {
	switch {
	case reflect.DeepEqual(w, d._fontList):
	case reflect.DeepEqual(w, d._fontName):
		d._fontName.SetText(d._fontList.GetSelection())
	}
}

func (d *FontDialogDirector) Showdialog() {}

type IWidget interface {
	Changed()
}

type ListBox struct {
	_director IDialogDirector
}

func (l *ListBox) Changed() {
	l._director.WidgetChanged(l)
}

func (l *ListBox) SetList() {}

func (l *ListBox) GetSelection() string {
	return ""
}

func NewListBox(d IDialogDirector) *ListBox {
	return &ListBox{
		_director: d,
	}
}

type EntryField struct {
	_director IDialogDirector
}

func (e *EntryField) Changed() {
	e._director.WidgetChanged(e)
}

func (e *EntryField) SetText(text string) {}

func NewEntryField(d IDialogDirector) *EntryField {
	return &EntryField{
		_director: d,
	}
}

type Button struct {
	_director IDialogDirector
}

func (b *Button) Changed() {
	b._director.WidgetChanged(b)
}

func (b *Button) SetText(text string) {}

func NewButton(d IDialogDirector) *Button {
	return &Button{
		_director: d,
	}
}
