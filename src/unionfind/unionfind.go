package main

import (
	"io"
	"os"
	"fmt"
	"log"
)

type UF interface {
	Union(p, q int)
	Find(p int) int
}

type QuickFindUF struct {
	ids		[]int	// access to component id
	count	int		// number of components
}

func NewQuickFindUF(N int) *QuickFindUF {
	if (N < 1) {
		panic(fmt.Sprintf("Invalid parameter N %d", N))
	}
	ids := make([]int, N)
	for i := 0; i < len(ids); i++ {
		ids[i] = i;
	}
	return &QuickFindUF{ids, N}
}

func (qf *QuickFindUF) Union(p, q int) {
	if qf.Connected(p, q) {
		return
	}
	pId := qf.ids[p]
	qId := qf.ids[q]
	for i, id := range qf.ids {
		if id == pId {
			qf.ids[i] = qId
		}
	}
	qf.count--
}
 
func (qf *QuickFindUF) Find(p int) int {
	return qf.ids[p]
} 

func (qf *QuickFindUF) Connected(p, q int) bool {
	return qf.Find(p) == qf.Find(q)
} 

func (qf *QuickFindUF) Count() int {
	return qf.count;
} 

type QuickUnionUF struct {
	ids		[]int	// access to component id
	count 	int		// number of components
}

func NewQuickUnionUF(N int) *QuickUnionUF {
	if N < 1 {
		panic(fmt.Sprintf("Invalid parameter N %d", N))
	}
	ids := make([]int, N)
	for i := 0; i < len(ids); i++ {
		ids[i] = i
	}
	count := N
	return &QuickUnionUF{ids, count}
}

func (qu *QuickUnionUF) Union(p, q int) {
	pRoot := qu.Find(p)
	qRoot := qu.Find(q)
	
	if pRoot == qRoot {
		return
	}
	
	qu.ids[pRoot] = qRoot
	qu.count--
}

func (qu *QuickUnionUF) Find(p int) int {
	for p != qu.ids[p] {
		p = qu.ids[p]
	}
	return p
}

func (qu *QuickUnionUF) Connected(p, q int) bool {
	pRoot := qu.Find(p)
	qRoot := qu.Find(q)
	return pRoot == qRoot
}

func (qu *QuickUnionUF) Count() int {
	return qu.count
}

// quick-union with path compression
type QuickUnionPathCompUF QuickUnionUF

func NewQuickUnionPathCompUF(N int) *QuickUnionPathCompUF {
	if N < 1 {
		panic(fmt.Sprintf("Invalid parameter N %d", N))
	}
	ids := make([]int, N)
	for i := 0; i < len(ids); i++ {
		ids[i] = i
	}
	count := N
	return &QuickUnionPathCompUF{ids, count}
}

func (qupc *QuickUnionPathCompUF) Union(p, q int) {
	pRoot := qupc.Find(p)
	qRoot := qupc.Find(q)
	
	if pRoot == qRoot {
		return
	}
	
	qupc.ids[pRoot] = qRoot
	qupc.count--
}

func (qupc *QuickUnionPathCompUF) Find(p int) int {
	n := p
	for p != qupc.ids[p] {
		p = qupc.ids[p]
	}
	// now p is the root node
	for n != qupc.ids[n] {
		t := qupc.ids[n]
		qupc.ids[n] = p	// make n links directly to the root
		n = t
	}
	return p
}

func (qupc *QuickUnionPathCompUF) Connected(p, q int) bool {
	pRoot := qupc.Find(p)
	qRoot := qupc.Find(q)
	return pRoot == qRoot
}

func (qupc *QuickUnionPathCompUF) Count() int {
	return qupc.count
}

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

type WeightedQuickUnionPathCompUF WeightedQuickUnionUF

func NewWeightedQuickUnionPathCompUF(N int) *WeightedQuickUnionPathCompUF {
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
	return &WeightedQuickUnionPathCompUF{ids, sz, N}
}

func (wqupc *WeightedQuickUnionPathCompUF) Union(p, q int) {
	pRoot := wqupc.Find(p)
	qRoot := wqupc.Find(q)
	if pRoot == qRoot {
		return
	}
	if wqupc.sz[pRoot] < wqupc.sz[qRoot] {
		wqupc.ids[pRoot] = qRoot
		wqupc.sz[qRoot] += wqupc.sz[pRoot]
	} else {
		wqupc.ids[qRoot] = pRoot
		wqupc.sz[pRoot] += wqupc.sz[qRoot]
	}
	wqupc.count--
}

func (wqupc *WeightedQuickUnionPathCompUF) Find(p int) int {
	root := p
	for root != wqupc.ids[root] {
		root = wqupc.ids[root]
	}
	for p != root {
		t := wqupc.ids[p]
		wqupc.ids[p] = root
		p = t
	}
	return root
}

func (wqupc *WeightedQuickUnionPathCompUF) Connected(p, q int) bool {
	pRoot := wqupc.Find(p)
	qRoot := wqupc.Find(q)
	return pRoot == qRoot
}

func (wqupc *WeightedQuickUnionPathCompUF) Count() int {
	return wqupc.count
}

func fetchFromFile(fname string, chnl chan<- int) {
	fh, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()
	
	var sz int
	_, err = fmt.Fscanf(fh, "%d\n", &sz)
	if err != nil {
		log.Fatalln(err)
	}
	chnl <- sz
	
	eof := false
	var p, q int
	for !eof {
		_, err := fmt.Fscanf(fh, "%d %d\n", &p, &q)
		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				log.Fatalln(err)
			}
		}
		chnl <- p
		chnl <- q
	}
	chnl <- -1	// means finished
}
var tinyFile string = "tinyUF.txt"
var mediumFile string = "mediumUF.txt"
var largeFile string = "largeUF.txt"
func main() {
	fmt.Println("--- Quick-Find for tinyUF.txt ---")
	tinyChan := make(chan int)
	go fetchFromFile(tinyFile, tinyChan)
	sz := <-tinyChan
	qfuf := NewQuickFindUF(sz)
	var p, q int
	for true {
		p = <-tinyChan
		if p == -1 {
			break
		}
		
		q = <-tinyChan
		if q == -1 {
			break
		}
		
		if !qfuf.Connected(p, q) {
			qfuf.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
		}
	}
	fmt.Printf("%d components\n", qfuf.Count())
	
	/*
	fmt.Println("--- Quick-Union with path compression for mediumUF.txt ---")
	medChan := make(chan int)
	go fetchFromFile(mediumFile, medChan)
	sz = <-medChan
	qupc := NewQuickUnionPathCompUF(sz)
	for true {
		p = <-medChan
		if p == -1 {
			break
		}
		
		q = <-medChan
		if q == -1 {
			break
		}
		if !qupc.Connected(p, q) {
			qupc.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
		}
	}
	fmt.Printf("%d components\n", qupc.Count())
	*/
	
	fmt.Println("--- Weighted Quick-Union for largeUF.txt ---")
	largeChan := make(chan int)
	go fetchFromFile(largeFile, largeChan)
	sz = <-largeChan
	wquf := NewWeightedQuickUnionPathCompUF(sz)
	for true {
		p = <-largeChan
		if p == -1 {
			break
		}
		
		q = <-largeChan
		if q == -1 {
			break
		}
		if !wquf.Connected(p, q) {
			wquf.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
		}
	}
	fmt.Printf("%d components\n", wquf.Count())
}