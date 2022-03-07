package iterator

type IIterator interface {
	First()
	Next()
	IsDone() bool
	CurrentItem() interface{}
}

type ListIterator struct {
	_current int
	_list    IList
}

func (l *ListIterator) First() {
	l._current = 0
}

func (l *ListIterator) Next() {
	l._current++
}

func (l *ListIterator) IsDone() bool {
	return l._current >= l._list.Count()
}

func (l *ListIterator) CurrentItem() interface{} {
	return l._list.Get(l._current)
}

func NewListIterator(l *List) *ListIterator {
	return &ListIterator{
		_current: 0,
		_list:    l,
	}
}

type IList interface {
	CreateIterator() IIterator
	Count() int
	Get(int) interface{}
}

type List struct{}

func (l *List) CreateIterator() IIterator {
	return NewListIterator(l)
}

func (l *List) Count() int {
	panic("not implemented") // TODO: Implement
}

func (l *List) Get(_ int) interface{} {
	panic("not implemented") // TODO: Implement
}

type SkipList struct{}

func (l *SkipList) CreateIterator() IIterator {
	panic("not implemented") // TODO: Implement
}

func (l *SkipList) Count() int {
	panic("not implemented") // TODO: Implement
}

func (l *SkipList) Get(_ int) interface{} {
	panic("not implemented") // TODO: Implement
}
