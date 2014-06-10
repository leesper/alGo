package cntr

import (
	"fmt"
)

type PQueue struct {
	keys    []interface{}
	length  int
	compare func(interface{}, interface{}) bool
}

func NewPQ(cmp func(interface{}, interface{}) bool) *PQueue {
	// arr[0] not use
	arr := make([]interface{}, 1)
	arr[0] = 0
	return &PQueue{
		keys:    arr,
		compare: cmp,
	}
}

func (pq *PQueue) Insert(v interface{}) {
	pq.keys = append(pq.keys, v)
	pq.length++
	pq.swim(pq.length)
}

func (pq *PQueue) swim(k int) {
	for k > 1 && pq.compare(pq.keys[k/2], pq.keys[k]) {
		pq.keys[k/2], pq.keys[k] = pq.keys[k], pq.keys[k/2]
		k = k / 2
	}
}

func (pq *PQueue) pq() interface{} {
	if pq.Empty() {
		panic(fmt.Sprintf("queue is empty"))
	}
	return pq.keys[1]
}

func (pq *PQueue) PrintOut() {
	fmt.Println(pq.keys[1:])
}

func (pq *PQueue) Del() interface{} {
	if pq.Empty() {
		panic(fmt.Sprintf("Trying to delete an empty queue"))
	}
	ret := pq.keys[1]
	pq.keys[1] = pq.keys[pq.length]
	pq.keys = pq.keys[:pq.length]
	pq.length--
	pq.sink(1)
	return ret
}

func (pq *PQueue) sink(k int) {
	for 2*k <= pq.length {
		j := 2 * k
		if j < pq.length && pq.compare(pq.keys[j], pq.keys[j+1]) {
			j++
		}
		if !pq.compare(pq.keys[k], pq.keys[j]) {
			break
		}
		pq.keys[k], pq.keys[j] = pq.keys[j], pq.keys[k]
		k = j
	}
}

func (pq *PQueue) IsHeapified() bool {
	k := 1
	for 2*k <= pq.length {
		j := 2 * k
		if pq.compare(pq.keys[k], pq.keys[j]) || (j < pq.length && pq.compare(pq.keys[k], pq.keys[j+1])) {
			return false
		}
		k++
	}
	return true
}

func (pq *PQueue) Empty() bool {
	return pq.length == 0
}

func (pq *PQueue) Size() int {
	return pq.length
}
