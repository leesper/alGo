package sorter_test

import (
	"fmt"
	"sorter"
	"strings"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (b ByAge) Length() int {
	return len(b)
}

func (b ByAge) Exchange(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func TestSelection(t *testing.T) {
	fmt.Println("====== Selection Sort ======")
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sorter.Selection(ByAge(people), 0, len(people))
	fmt.Println(people)
}

type Planet struct {
	name string
	mass float64
	dist float64
}

func (p *Planet) String() string {
	return fmt.Sprintf("%s mass: %.3f distance: %.2f", p.name, p.mass, p.dist)
}

func TestMultiKeySorter(t *testing.T) {
	fmt.Println("====== Multi Key Sorter ======")
	planets := []Planet{
		{"Mercury", 0.055, 0.4},
		{"Venus", 0.815, 0.7},
		{"Earth", 1.0, 1.0},
		{"Mars", 0.107, 1.5},
	}

	name := func(o1, o2 interface{}) bool {
		p1 := o1.(Planet)
		p2 := o2.(Planet)
		return strings.ToLower(p1.name) < strings.ToLower(p2.name)
	}

	distance := func(o1, o2 interface{}) bool {
		p1 := o1.(Planet)
		p2 := o2.(Planet)
		return p1.dist < p2.dist
	}

	mass := func(o1, o2 interface{}) bool {
		p1 := o1.(Planet)
		p2 := o2.(Planet)
		return p1.mass < p2.mass
	}

	sorter.By(name).Sort(planets)
	fmt.Println("By name: ", planets)

	sorter.By(distance).Sort(planets)
	fmt.Println("By distance: ", planets)

	sorter.By(mass).Sort(planets)
	fmt.Println("By mass: ", planets)
}

func TestInsertion(t *testing.T) {
	fmt.Println("====== Insertion Sort ======")
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sorter.Insertion(ByAge(people), 0, len(people))
	fmt.Println(people)
}

func TestShell(t *testing.T) {
	fmt.Println("====== Shell Sort ======")
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sorter.Shell(ByAge(people), 0, len(people))
	fmt.Println(people)
}

func TestMergeSort(t *testing.T) {
	fmt.Println("====== Merge Sort ======")
	ints := []int{41, 58, 72, 18, 90, 88, 77, 26, 98, 93}
	sorter.MergeSort(ints)
	fmt.Println(ints)
}

func TestMergeSortBU(t *testing.T) {
	fmt.Println("====== Merge Sort ======")
	ints := []int{41, 58, 72, 18, 90, 88, 77, 26, 98, 93}
	sorter.MergeSortBU(ints)
	fmt.Println(ints)
}

func TestQuickSort(t *testing.T) {
	fmt.Println("====== Quick Sort ======")
	ints := []int{41, 58, 72, 18, 90, 88, 77, 26, 98, 93}
	sorter.QuickSort(ints)
	fmt.Println(ints)
}

func TestHeapSort(t *testing.T) {
	fmt.Println("====== Heap Sort ======")
	ints := []int{41, 58, 72, 18, 90, 88, 77, 26, 98, 93}
	sorter.HeapSort(ints)
	fmt.Println(ints)
}
