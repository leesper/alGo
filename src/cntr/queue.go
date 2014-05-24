package cntr

import (
	"fmt"
)

type Queue struct {
	first	*node
	last	*node
	n		int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue)	Enqueue(item interface{}) {
	if item == nil {
		panic(fmt.Sprintf("enqueue nil items"))
	}
	oldlast := q.last
	nd := newNode(item)
	q.last = nd
	if q.Empty() {
		q.first = q.last
	} else {
		oldlast.next = q.last
	}
	q.n++
}

func (q *Queue) Dequeue() interface{} {
	if q.Empty() {
		panic(fmt.Sprintf("dequeue empty queue"))
	}
	item := q.first.item
	q.first = q.first.next
	if q.Empty() {
		q.last = nil
	}
	q.n--
	return item
}

func (q *Queue) Empty() bool {
	return q.first == nil
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) Iterator() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for iter := q.first; iter != nil; iter = iter.next {
			ch <- iter.item
		}
		close(ch)
	}()
	return ch
}