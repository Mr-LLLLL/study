package base

type Application struct{}

func (Application) Add(*Document) {}

type Document struct{}

func (Document) Open() {}

func (Document) Paste() {}

func NewDocument() *Document {
	return new(Document)
}
