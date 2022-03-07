package list

import "iterator/iterator"

type IList interface {
	CreateIterator() iterator.IIterator
}

type List struct {
}
