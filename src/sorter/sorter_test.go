package sorter_test

import (
	"fmt"
	"sorter"
	"strings"
    "testing"
)

type Person struct {
	Name	string
	Age		int
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
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sorter.Selection(ByAge(people))
	fmt.Println(people)
}

type Planet struct {
	name	string
	mass	float64
	dist	float64
}

func (p *Planet) String() string {
	return fmt.Sprintf("%s mass: %.3f distance: %.2f", p.name, p.mass, p.dist)
}

func TestMultiKeySorter(t *testing.T) {
	planets := []Planet {
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