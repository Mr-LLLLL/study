package pq

import "errors"

type Pq struct {
	arr  []int
	size int
}

func (q *Pq) GetMax() (int, error) {
	if q.size == 0 {
		return 0, errors.New("Priority queue is empty")
	}
	return q.arr[0], nil
}

func (q *Pq) Insert(e int) {
	q.arr = append(q.arr, e)
	q.percolateUp(q.size)
	q.size++
}

func (q *Pq) DelMax() int {
	maxElem := q.arr[0]
	q.arr[0] = q.arr[q.size-1]
	q.percolateDown(0)
	q.size--
	q.arr = q.arr[:q.size]
	return maxElem
}

func (q *Pq) Size() int {
	return q.size
}

func (q *Pq) Empty() bool {
	return q.size == 0
}

func (q *Pq) percolateUp(i int) int {
	for q.parentValid(i) {
		j := q.parent(i)
		if q.arr[i] <= q.arr[j] {
			break
		}
		q.arr[i], q.arr[j] = q.arr[j], q.arr[i]
		i = j
	}
	return i
}

func (q *Pq) percolateDown(i int) int {
	for true {
		j := q.properParent(i)
		if i == j {
			break
		}
		q.arr[i], q.arr[j] = q.arr[j], q.arr[i]
		i = j
	}
	return i
}

func (q *Pq) inHeap(i int) bool {
	return -1 < i && i < q.size
}

func (q *Pq) parent(i int) int {
	return (i - 1) >> 1
}

func (q *Pq) lastInternal() int {
	return q.parent(q.size - 1)
}

func (q *Pq) lChild(i int) int {
	return 1 + (i << 1)
}

func (q *Pq) rChild(i int) int {
	return (1 + i) << 1
}

func (q *Pq) parentValid(i int) bool {
	return i > 0
}

func (q *Pq) lChildValid(i int) bool {
	return q.inHeap(q.lChild(i))
}

func (q *Pq) rChildValid(i int) bool {
	return q.inHeap(q.rChild(i))
}

func (q *Pq) max(i, j int) int {
	if q.arr[i] >= q.arr[j] {
		return i
	}
	return j
}

func (q *Pq) properParent(i int) int {
	var tmp int
	if q.lChildValid(i) {
		if q.rChildValid(i) {
			tmp = q.max(q.lChild(i), q.rChild(i))
		} else {
			tmp = q.lChild(i)
		}
	} else {
		if q.rChildValid(i) {
			tmp = q.rChild(i)
		} else {
			return i
		}
	}

	return q.max(i, tmp)
}

func (q *Pq) heapify() {
	for i := q.lastInternal(); q.inHeap(i); i-- {
		q.percolateDown(i)
	}
}

func NewPq(arr ...int) *Pq {
	dup := make([]int, 0)
	copy(dup, arr)
	q := &Pq{
		arr:  dup,
		size: len(dup),
	}
	q.heapify()
	return q
}

func HeapSort(arr []int, lo, hi int) {
	q := NewPq(arr[lo:hi]...)
	for q.Empty() {
		hi--
		arr[hi] = q.DelMax()
	}
}
