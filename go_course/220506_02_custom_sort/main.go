// https://www.youtube.com/watch?v=0E-q22d3QD4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=118&ab_channel=AprendaGo

package main

import (
	"fmt"
	"sort"
)

// type: keyword
// vehicle: name of the new type
// struct is the format of this new type
type Vehicle struct {
	Name       string
	Power      int
	Efficiency int
}

// how is this related to C typedef?

type ByPower []Vehicle

func (arr ByPower) Len() int {
	return len(arr)
}

func (arr ByPower) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr ByPower) Less(i, j int) bool {
	return arr[i].Power < arr[j].Power
}

func main() {
	cars := []Vehicle{
		{Name: "Gol", Power: 10, Efficiency: 1},
		{Name: "Fiesta", Power: 9, Efficiency: 2},
		{Name: "Fox", Power: 11, Efficiency: 0},
		{Name: "Corolla", Power: 30, Efficiency: -8},
		{Name: "Sandero", Power: 3, Efficiency: 1},
	}

	fmt.Println("original slice: ", cars)

	sort.Sort(ByPower(cars))

	fmt.Println("sorted by power slice: ", cars)
}

// an interface is a type that we define
// it is a set of functions with a given name
