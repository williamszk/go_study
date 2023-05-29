package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func main() {

	p1 := person{"James", 32}
	p2 := person{"Monkey", 31}
	p3 := person{"Bob", 30}
	p4 := person{"Alice", 40}

	people := []person{p1, p2, p3, p4}

	fmt.Println("before:", people)

	// sort the people slice using first and using age
	// create a type that will group the methods needed for sorting
	// create byAge and byName which will need to have some methods
	sort.Sort(byAge(people))
	fmt.Println("sort by age:", people)

	sort.Sort(byName(people))
	fmt.Println("sort by name:", people)

}

// we say that the underlying type of byAge is []person
type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

type byName []person

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].first < a[j].first }
