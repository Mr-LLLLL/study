package scanner

import (
	"facade/base"
)

type IScanner interface {
	Scan()
}

type Scanner struct {
	_inputStream *base.IsStream
}

func (s *Scanner) Scan() {}

func NewScanner(input *base.IsStream) *Scanner {
	return &Scanner{
		_inputStream: input,
	}
}
