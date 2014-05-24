package cntr_test

import (
	"fmt"
	"cntr"
	"testing"
)

func TestStack(t *testing.T) {
	fmt.Println("========Testing For Stack========")
	stk := cntr.NewStack()
	stk.Push("My")
	stk.Push("Favourite")
	stk.Push("Programming Language")
	stk.Push("Is")
	stk.Push("Golang")
	
	for item := range stk.Iterator() {
		fmt.Println(item.(string))
	}
	
	fmt.Printf("stack size: %d\n", stk.Size())
	for !stk.Empty() {
		fmt.Println(stk.Pop())
	}
	fmt.Println(stk.Empty())
}

func TestQueue(t *testing.T) {
	fmt.Println("========Testing For Queue========")
	que := cntr.NewQueue()
	que.Enqueue("MY")
	que.Enqueue("FAVOURITE")
	que.Enqueue("PROGRAMMING LANGUAGE")
	que.Enqueue("IS")
	que.Enqueue("GOLANG")
	fmt.Println(que.Size())
	
	for item := range que.Iterator() {
		fmt.Println(item)
	}
	
	for !que.Empty() {
		fmt.Println(que.Dequeue())
	}
}