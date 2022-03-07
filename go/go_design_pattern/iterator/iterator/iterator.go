package iterator

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

type ArrayInt []int

func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: ArrayInt{},
		index:    0,
	}
}

type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (it *ArrayIntIterator) HasNext() bool {
	return it.index < len(it.arrayInt)
}

func (it *ArrayIntIterator) Next() {
	it.index++
}

func (it *ArrayIntIterator) CurrentItem() interface{} {
	return it.arrayInt[it.index]
}

