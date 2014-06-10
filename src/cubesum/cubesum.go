package main

import (
	"cntr"
	"fmt"
	"math"
)

type Tuple struct {
	csum int64
	a    int64
	b    int64
}

func (t *Tuple) String() string {
	return fmt.Sprintf("(%d, %d, %d)", t.csum, t.a, t.b)
}

func NewTuple(a, b int64) *Tuple {
	cubesum := int64(math.Pow(float64(a), float64(3)) + math.Pow(float64(b), float64(3)))
	return &Tuple{
		csum: cubesum,
		a:    a,
		b:    b,
	}
}

func main() {
	var N int64 = 12
	minpq := cntr.NewPQ(func(m, n interface{}) bool {
		t1, t2 := m.(*Tuple), n.(*Tuple)
		if t1.csum > t2.csum {
			return true
		} else if t1.csum == t2.csum {
			if t1.a > t2.a {
				return true
			} else if t1.a == t2.a {
				if t1.b > t2.b {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		}
		return false
	})

	var i int64
	for i = 0; i <= N; i++ {
		minpq.Insert(NewTuple(i, 0))
	}
	//minpq.PrintOut()
	var prev *Tuple = nil
	var printed bool = false
	for !minpq.Empty() {
		curr := minpq.Del().(*Tuple)
		//fmt.Printf("%d^3 + %d^3 = %d\n", curr.a, curr.b, curr.csum)
		if prev != nil && prev.csum == curr.csum {
			if !printed && prev.a != curr.a && prev.a != curr.b && prev.b != curr.a && prev.b != curr.b {
				fmt.Printf("%d^3 + %d^3 = %d^3 + %d^3\n", prev.a, prev.b, curr.a, curr.b)
				printed = true
			}
		} else {
			printed = false
		}
		if curr.b < N {
			minpq.Insert(NewTuple(curr.a, curr.b+1))
		}
		prev = curr
	}
}
