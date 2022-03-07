package stack

import (
	"errors"
)

type stack []interface{}

func (s *stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) Pop() (res interface{}, err error) {
	if len(*s) == 0 {
		return nil, errors.New("stack is empty")
	}
	res = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *stack) Push(e interface{}) (res interface{}) {
	res = e
	*s = append(*s, e)
	return
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func (s *stack) Size() int {
	return len(*s)
}

func New() *stack {
	return &stack{}
}
