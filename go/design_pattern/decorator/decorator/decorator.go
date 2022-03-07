package decorator

type IVisualComponent interface {
	Draw()
	Resize()
}

type TextView struct{}

func (t *TextView) Draw() {}

func (t *TextView) Resize() {}

type Decorator struct {
	_component IVisualComponent
}

func (d *Decorator) Draw() {
	d._component.Draw()
}

func (d *Decorator) Resize() {
	d._component.Resize()
}

func NewDecorator(i IVisualComponent) *Decorator {
	return &Decorator{
		_component: i,
	}
}

type BorderDecorator struct {
	_component IVisualComponent
	_width     int
}

func (b *BorderDecorator) Draw() {
	b._component.Draw()
	b.DrawBorder()
}

func (b *BorderDecorator) Resize() {
	panic("not implemented") // TODO: Implement
}

func (b *BorderDecorator) DrawBorder() {
}

func NewBorderDecorator(i IVisualComponent, w int) *BorderDecorator {
	return &BorderDecorator{
		_component: i,
		_width:     w,
	}
}
