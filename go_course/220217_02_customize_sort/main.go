// https://www.youtube.com/watch?v=0E-q22d3QD4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=118&ab_channel=AprendaGo

package main

import (
	"fmt"
	"sort"
)

type vehicle struct {
	name       string
	horsePower int
	efficiency float64
}

// create new types which are slice of vehicles
type orderByHorsePower []vehicle

func (sliVehic orderByHorsePower) Len() int {
	return len(sliVehic)
}

func (sliVehic orderByHorsePower) Less(i, j int) bool {
	return sliVehic[i].horsePower < sliVehic[j].horsePower
}

func (sliVehic orderByHorsePower) Swap(i, j int) {
	sliVehic[i], sliVehic[j] = sliVehic[j], sliVehic[i]
}

type orderByEfficiency []vehicle

func (sliVehic orderByEfficiency) Len() int {
	return len(sliVehic)
}

func (sliVehic orderByEfficiency) Less(i, j int) bool {
	return sliVehic[i].efficiency > sliVehic[j].efficiency
}

func (sliVehic orderByEfficiency) Swap(i, j int) {
	sliVehic[i], sliVehic[j] = sliVehic[j], sliVehic[i]
}

type orderByCost []vehicle

func (sliVehic orderByCost) Len() int {
	return len(sliVehic)
}

func (sliVehic orderByCost) Less(i, j int) bool {
	return sliVehic[i].efficiency < sliVehic[j].efficiency
}

func (sliVehic orderByCost) Swap(i, j int) {
	sliVehic[i], sliVehic[j] = sliVehic[j], sliVehic[i]
}

func main() {

	sliVehic := []vehicle{
		{
			name:       "Lancer",
			horsePower: 3000,
			efficiency: 4,
		},
		{
			name:       "Gol",
			horsePower: 1000,
			efficiency: 2,
		},
		{
			name:       "Uno",
			horsePower: 500,
			efficiency: 3,
		},
	}

	fmt.Println(sliVehic) // [{Lancer 3000 4} {Gol 1000 2} {Uno 500 3}]

	sort.Sort(orderByHorsePower(sliVehic))

	fmt.Println(sliVehic) // [{Uno 500 3} {Gol 1000 2} {Lancer 3000 4}]

	sort.Sort(orderByEfficiency(sliVehic))

	fmt.Println(sliVehic) // [{Lancer 3000 4} {Uno 500 3} {Gol 1000 2}]

	sort.Sort(orderByCost(sliVehic))

	fmt.Println(sliVehic) // [{Gol 1000 2} {Uno 500 3} {Lancer 3000 4}]

}

// next
// https://www.youtube.com/watch?v=4vCb7jmwkzM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=119&ab_channel=AprendaGo
