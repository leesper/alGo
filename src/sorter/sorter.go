package sorter

import (
	"reflect"
)

type Sortable interface {
	Length()			int
	Less(i, j int)		bool
	Exchange(i, j int)
}

type MultiKeySorter struct {
	coll	interface{}
	lesser	func(o1, o2 interface{}) bool
}

func (mks *MultiKeySorter) Length() int {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		return slice.Len()
	}
	panic("passing a non-slice type")
}

func (mks *MultiKeySorter) Exchange(i, j int) {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		temp := reflect.ValueOf(slice.Index(i).Interface())
		slice.Index(i).Set(reflect.ValueOf(slice.Index(j).Interface()))
		slice.Index(j).Set(temp)
		return
	}
	panic("passing a non-slice type")
}

func (mks *MultiKeySorter) Less(i, j int) bool {
	if reflect.TypeOf(mks.coll).Kind() == reflect.Slice {
		slice := reflect.ValueOf(mks.coll)
		return mks.lesser(slice.Index(i).Interface(), slice.Index(j).Interface())
	}
	panic("passing a non-slice type")
}

type By func(o1, o2 interface{}) bool

func (by By) SelectionSort(coll interface{}) {
	mks := &MultiKeySorter {
		coll:		coll,
		lesser:		by,
	}
	Selection(mks)
}

func Selection(coll Sortable) {
	N := coll.Length();
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if coll.Less(j, min) {
				min = j
			}
		}
		coll.Exchange(i, min)
	}
}

