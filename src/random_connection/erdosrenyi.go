package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"random_connection/wquf"
	"strconv"
	"math/rand"
)

func usage() string {
	return fmt.Sprintf("%s N", os.Args[0])
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalln(usage())
	}
	
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	
	if N <= 1 {
		log.Fatalln("error, N must be greater than 1")
	}
	
	wqu := wquf.NewWeightedQuickUnionUF(N)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0
	for wqu.Count() > 1 {
		rp := rng.Intn(N)
		rq := rng.Intn(N)
		if !wqu.Connected(rp, rq) {
			wqu.Union(rp, rq)
			count++
		}
	}
	fmt.Printf("number of connections: %d\n", count)
}