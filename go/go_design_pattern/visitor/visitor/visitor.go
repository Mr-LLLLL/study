package visitor

import (
	"fmt"
	"path"
)

type Visitor interface {
	Visit(IResourceFile) error
}

type IResourceFile interface {
	Accept(Visitor) error
}

func NewResourceFile(filepath string) (IResourceFile, error) {
	switch path.Ext(filepath) {
	case ".ppt":
		return &PPTFile{
			path: filepath,
		}, nil
	case ".pdf":
		return &PdfFile{
			path: filepath,
		}, nil
	default:
		return nil, fmt.Errorf("not found file type:%s", filepath)
	}
}

type PdfFile struct {
	path string
}

func (f *PdfFile) Accept(visitor Visitor) error {
	return visitor.Visit(f)
}

type PPTFile struct {
	path string
}

func (f *PPTFile) Accept(visitor Visitor) error {
	return visitor.Visit(f)
}

type Compressor struct{}

func (c *Compressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PPTFile:
		return c.VisitPPTFile(f)
	case *PdfFile:
		return c.VisitPDFFile(f)
	default:
		return fmt.Errorf("not found resource type: %#v", r)
	}
}

func (c *Compressor) VisitPPTFile(f *PPTFile) error {
	fmt.Println("this is ppt file")
	return nil
}
func (c *Compressor) VisitPDFFile(f *PdfFile) error {
	fmt.Println("this is pdf file")
	return nil
}
