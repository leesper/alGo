package main

import (
	"fmt"
)

type WeightedQuickUnionFindMax struct {
	ids		[]int
	sz		[]int
	max		[]int
	count	int
}

func NewWeightedQuickUnionFindMax(N int) *WeightedQuickUnionFindMax {
	if N < 1 {
		panic(fmt.Sprintf("Invalid parameter N: %d", N))
	}
	ids := make([]int, N)
	sz := make([]int, N)
	max := make([]int, N)
	count := N
	
	for i := 0; i < N; i++ {
		ids[i] = i
		max[i] = i
		sz[i] = 1
	}
	return &WeightedQuickUnionFindMax{ids, sz, max, count}
}

func (wqufm *WeightedQuickUnionFindMax) Find(p int) int {
	for p != wqufm.ids[p] {
		p = wqufm.ids[p]
	}
	return wqufm.max[p]
}

func (wqufm *WeightedQuickUnionFindMax) Union(p, q int) {
	pRoot := wqufm.Find(p)
	qRoot := wqufm.Find(q)
	if pRoot == qRoot {
		return
	}
	
	if wqufm.sz[pRoot] < wqufm.sz[qRoot] {
		wqufm.ids[pRoot] = qRoot
		wqufm.sz[qRoot] += wqufm.sz[pRoot]
		if wqufm.max[qRoot] < wqufm.max[pRoot] {
			wqufm.max[qRoot] = wqufm.max[pRoot]
		}
	} else {
		wqufm.ids[qRoot] = pRoot
		wqufm.sz[pRoot] += wqufm.sz[qRoot]
		if wqufm.max[pRoot] < wqufm.max[qRoot] {
			wqufm.max[pRoot] = wqufm.max[qRoot]
		}
	}
	wqufm.count--
}

func (wqufm *WeightedQuickUnionFindMax) Connected(p, q int) bool {
	return wqufm.Find(p) == wqufm.Find(q)
}

func (wqufm *WeightedQuickUnionFindMax) Count() int {
	return wqufm.count
}

type RemoveSet struct {
	wqufm		*WeightedQuickUnionFindMax
	count		int
}

func NewRemoveSet(N int) *RemoveSet {
	wqu := NewWeightedQuickUnionFindMax(N)
	return &RemoveSet{wqu, N}
}

func (rs *RemoveSet) Remove(p int) bool {
	if p < 0 || p >= rs.count - 1 {
		return false
	}
	rs.wqufm.Union(p + 1, p)
	return true
}

func (rs *RemoveSet) Successor(p int) (int, bool) {
	if p < 0 || p >= rs.count - 1 {
		return p, false
	}
	return rs.wqufm.Find(p), true
}

func main() {
	rs := NewRemoveSet(9)
	
	rs.Remove(1)
	if su, ok := rs.Successor(1); ok {
		fmt.Println(su)
	}
	
	rs.Remove(2)
	if su, ok := rs.Successor(2); ok {
		fmt.Println(su)
	}
}