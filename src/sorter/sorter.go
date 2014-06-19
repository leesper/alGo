// Package sorter implements a series of sorting algorithms
// based on the book "Algorithms 4ed" by Sedgewick
// Author: Leesper
// Email: pascal7718@gmail.com 394683518@qq.com

package sorter

import (
	"math/rand"
	"reflect"
	"time"
)

// Sortable is an interface for a collection to be sortable
// All sorting algorithms sort collection that is a kind of Sortable
type Sortable interface {
	Length() int
	Less(i, j int) bool
	Exchange(i, j int)
}

// multiKeySorter is a kind of Sortable used for sorting according to different keys
// One can write different lesser functions for a user-defined type and call
// By(*lesser*).Sort(*coll*) to sort the collection
type multiKeySorter struct {
	coll   interface{}
	lesser func(o1, o2 interface{}) bool
}

func (mks *multiKeySorter) Length() int {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		return slice.Len()
	}
	panic("passing a non-slice type")
}

func (mks *multiKeySorter) Exchange(i, j int) {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		temp := reflect.ValueOf(slice.Index(i).Interface())
		slice.Index(i).Set(reflect.ValueOf(slice.Index(j).Interface()))
		slice.Index(j).Set(temp)
		return
	}
	panic("passing a non-slice type")
}

func (mks *multiKeySorter) Less(i, j int) bool {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		return mks.lesser(slice.Index(i).Interface(), slice.Index(j).Interface())
	}
	panic("passing a non-slice type")
}

// By is a function type for multiple-key sorting
type By func(o1, o2 interface{}) bool

// Sort sorts the slice by lesser func passing in
func (by By) Sort(coll interface{}) {
	mks := &multiKeySorter{
		coll:   coll,
		lesser: by,
	}
	Shell(mks, 0, mks.Length())
}

// Selection sorts coll[bgn..end)
func Selection(coll Sortable, bgn, end int) {
	for i := bgn; i < end; i++ {
		min := i
		for j := i + 1; j < end && coll.Less(j, min); j++ {
			min = j
		}
		coll.Exchange(i, min)
	}
}

// Insertion sorts coll[bgn..end)
func Insertion(coll Sortable, bgn, end int) {
	for i := bgn; i < end; i++ {
		for j := i; j > bgn && coll.Less(j, j-1); j-- {
			coll.Exchange(j, j-1)
		}
	}
}

/* The original shell sort proto type
func Shell(coll Sortable) {
	N := coll.Length()
	h := 1
	for h < N / 3 {
		h = 3 * h + 1	// making incremental sequences
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h; j -= h {
				if coll.Less(j, j - h) {
					coll.Exchange(j, j - h)
				}
			}
		}
		h = h / 3
	}
}
*/

// Shell sorts coll[bgn..end)
func Shell(coll Sortable, bgn, end int) {
	h := 1
	for h < (end-bgn)/3 {
		h = 3*h + 1 // making incremental sequences
	}
	for h >= 1 {
		for i := h; i < (end - bgn); i++ {
			for j := i; j >= h && coll.Less(j, j-h); j -= h {
				coll.Exchange(j, j-h)
			}
		}
		h = h / 3
	}
}

// Example of merge sort for int slice
func MergeSort(arr []int) {
	aux := make([]int, len(arr))
	mergeSort(arr, aux, 0, len(arr)-1)
}

func mergeSort(arr, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(arr, aux, lo, mid)
	mergeSort(arr, aux, mid+1, hi)
	merge(arr, aux, lo, mid, hi)
}

// Example of merge sort in bottom-up style
func MergeSortBU(arr []int) {
	aux := make([]int, len(arr))
	N := len(arr)
	for sz := 1; sz < N; sz = sz + sz {
		for lo := 0; lo < N-sz; lo += sz + sz {
			merge(arr, aux, lo, lo+sz-1, min(lo+sz+sz-1, N-1))
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func merge(arr, aux []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		aux[i] = arr[i]
	}
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > hi {
			arr[k] = aux[i]
			i++
		} else if less(aux[i], aux[j]) {
			arr[k] = aux[i]
			i++
		} else {
			arr[k] = aux[j]
			j++
		}
	}
}

func less(a, b int) bool {
	return a-b < 0
}

func QuickSort(arr []int) {
	shuffle(arr)
	//quickSort(arr, 0, len(arr) - 1)
	quickSort3Way(arr, 0, len(arr)-1)
}

func quickSort3Way(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	v, lt, i, gt := arr[lo], lo, lo+1, hi
	for i <= gt {
		if arr[i] < v {
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}
	quickSort3Way(arr, lo, lt-1)
	quickSort3Way(arr, gt+1, hi)
}

func quickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(arr, lo, hi)
	quickSort(arr, lo, j-1)
	quickSort(arr, j+1, hi)
}

func partition(arr []int, lo, hi int) int {
	v, i, j := arr[lo], lo+1, hi
	for {
		for less(arr[i], v) {
			if i == hi {
				break
			}
			i++
		}
		for less(v, arr[j]) {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[j], arr[lo] = arr[lo], arr[j]
	return j
}

func shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	N := len(arr)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		arr[i], arr[r] = arr[r], arr[i]
	}
}

func HeapSort(arr []int) {
	N := len(arr)
	for k := N / 2; k >= 1; k-- {
		sink(arr, k, N)
	}
	for N > 1 {
		excher(arr, 1, N)
		N--
		sink(arr, 1, N)
	}
}

func sink(arr []int, k, N int) {
	for 2*k <= N {
		j := 2 * k
		if j < N && lesser(arr, j, j+1) {
			j++
		}
		if !lesser(arr, k, j) {
			break
		}
		excher(arr, k, j)
		k = j
	}
}

func lesser(arr []int, i, j int) bool {
	return arr[i-1] < arr[j-1]
}

func excher(arr []int, i, j int) {
	arr[i-1], arr[j-1] = arr[j-1], arr[i-1]
}
