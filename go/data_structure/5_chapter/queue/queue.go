package queue

import (
	"errors"
)

type node struct {
	data interface{}
	next *node
}

type queue struct {
	head, rear *node
	size       int
}

func (q *queue) Front() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.head.data, nil
}

func (q *queue) Back() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.rear.data, nil
}

func (q *queue) Enqueue(data interface{}) {
	n := &node{
		data: data,
		next: nil,
	}
	if q.size == 0 {
		q.rear = n
		q.head = n
	} else {
		q.rear.next = n
		q.rear = n
	}
	q.size++
}

func (q *queue) Dequeue() (data interface{}, err error) {
	if q.size == 0 {
		return nil, errors.New("no data any more")
	}
	node := q.head
	q.head = node.next
	q.size--

	if q.size == 0 {
		q.head = nil
		q.rear = nil
	}
	return node.data, nil
}

func (q *queue) Size() int {
	return q.size
}

func New() *queue {
	return &queue{
		head: nil,
		rear: nil,
		size: 0,
	}
}
