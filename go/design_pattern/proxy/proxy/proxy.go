package proxy

import "proxy/base"

type IGraphic interface {
	Draw(at *base.Point)
	HandleMouse(event *base.Event)

	GetExtent() *base.Point

	Load(from *base.Isstream)
	Save(to *base.Osstream)
}

type Image struct{}

func (i *Image) Draw(at *base.Point) {
	panic("not implemented") // TODO: Implement
}

func (i *Image) HandleMouse(event *base.Event) {
	panic("not implemented") // TODO: Implement
}

func (i *Image) GetExtent() *base.Point {
	panic("not implemented") // TODO: Implement
}

func (i *Image) Load(from *base.Isstream) {
	panic("not implemented") // TODO: Implement
}

func (i *Image) Save(to *base.Osstream) {
	panic("not implemented") // TODO: Implement
}

func NewImage(fileName string) *Image {
	return new(Image)
}

type ImageProxy struct {
	_image    *Image
	_extent   *base.Point
	_fileName string
}

func (i *ImageProxy) Draw(at *base.Point) {
	i.GetImage().Draw(at)
}

func (i *ImageProxy) HandleMouse(event *base.Event) {
	i.GetImage().HandleMouse(event)
}

func (i *ImageProxy) GetExtent() *base.Point {
	if i._extent == nil {
		i._extent = i.GetImage().GetExtent()
	}
	return i._extent
}

func (i *ImageProxy) Load(from *base.Isstream) {
	panic("not implemented") // TODO: Implement
}

func (i *ImageProxy) Save(to *base.Osstream) {
	panic("not implemented") // TODO: Implement
}

func (i *ImageProxy) GetImage() *Image {
	if i._image == nil {
		i._image = NewImage(i._fileName)
	}
	return i._image
}

func NewImageProxy(fileName string) *ImageProxy {
	return &ImageProxy{
		_fileName: fileName,
	}
}
