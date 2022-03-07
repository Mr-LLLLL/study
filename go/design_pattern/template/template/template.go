package template

type IView interface {
	DoDisplay()
}

type View struct {
	IView
}

func (v *View) Display() {
	v.SetFocus()
	v.DoDisplay()
	v.ResetFocus()
}

func (v *View) SetFocus() {}

func (v *View) ResetFocus() {}

type MyView struct {
	*View
}

func (m *MyView) DoDisplay() {
	// render the view's contents
}

func NewMyView() *MyView {
	myview := new(MyView)
	myview.View = &View{
		IView: myview,
	}
	return myview
}
