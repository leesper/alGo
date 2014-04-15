package wquf

import (
	"fmt"
)

type WeightedQuickUnionUF struct {
	ids		[]int
	sz		[]int
	count	int
}

func NewWeightedQuickUnionUF(N int) *WeightedQuickUnionUF {
	if N < 1 {
		fmt.Sprintf("Invalid parameter N: %d", N)
	}
	ids := make([]int, N)
	for i := 0; i < len(ids); i++ {
		ids[i] = i
	}
	sz := make([]int, N)
	for i := 0; i < len(sz); i++ {
		sz[i] = 1
	}
	return &WeightedQuickUnionUF{ids, sz, N}
}

func (wquf *WeightedQuickUnionUF) Union(p, q int) {
	pRoot := wquf.Find(p)
	qRoot := wquf.Find(q)
	if pRoot == qRoot {
		return
	}
	if wquf.sz[pRoot] < wquf.sz[qRoot] {
		wquf.ids[pRoot] = qRoot
		wquf.sz[qRoot] += wquf.sz[pRoot]
	} else {
		wquf.ids[qRoot] = pRoot
		wquf.sz[pRoot] += wquf.sz[qRoot]
	}
	wquf.count--
}

func (wquf *WeightedQuickUnionUF) Find(p int) int {
	for p != wquf.ids[p] {
		p = wquf.ids[p]
	}
	return p
}

func (wquf *WeightedQuickUnionUF) Connected(p, q int) bool {
	pRoot := wquf.Find(p)
	qRoot := wquf.Find(q)
	return pRoot == qRoot
}

func (wquf *WeightedQuickUnionUF) Count() int {
	return wquf.count
}