package cntr

import (
	"fmt"
)

type Stack struct {
	first	*node
	n		int
}

func NewStack() *Stack {
	return &Stack{}
}

func (stk *Stack) Push(item interface{}) {
	if item == nil {
		panic(fmt.Sprintf("pushing nil items"))
	}
	nd := newNode(item)
	oldfirst := stk.first
	stk.first = nd
	nd.next = oldfirst
	stk.n++
}

func (stk *Stack) Pop() interface{} {
	if stk.Empty() {
		panic(fmt.Sprintf("poping empty stack"))
	}
	item := stk.first.item
	stk.first = stk.first.next
	stk.n--
	return item
}

func (stk *Stack) Empty() bool {
	return stk.first == nil
}

func (stk *Stack) Size() int {
	return stk.n
}

func (stk *Stack) Iterator() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for it := stk.first; it != nil; it = it.next {
			ch <- it.item
		}
		close(ch)
	}()
	return ch
}