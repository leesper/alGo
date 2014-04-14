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
//var largeFile string = "largeUF.txt"
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
	
	fmt.Println("--- Quick-Union for mediumUF.txt ---")
	medChan := make(chan int)
	go fetchFromFile(mediumFile, medChan)
	sz = <-medChan
	quuf := NewQuickUnionUF(sz)
	for true {
		p = <-medChan
		if p == -1 {
			break
		}
		
		q = <-medChan
		if q == -1 {
			break
		}
		if !quuf.Connected(p, q) {
			quuf.Union(p, q)
			fmt.Printf("%d %d\n", p, q)
		}
	}
	fmt.Printf("%d components\n", quuf.Count())
}