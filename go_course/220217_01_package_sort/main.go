// https://www.youtube.com/watch?v=b67JIGYM6Hc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=117&ab_channel=AprendaGo

package main

import (
	"fmt"
	"sort"
)

func main() {

	myslice := []string{"A", "b", "C", "c", "D", "atirei", "10"}

	sort.Strings(myslice)

	fmt.Println(myslice)

	myslice1 := []int{10, 2, 3, 9, 5, 6, 8}
	sort.Ints(myslice1)

	fmt.Println(myslice1)

	myslice2 := []float64{10, 3, 50, 2, 8, 9}
	sort.Float64s(myslice2)

	fmt.Println(myslice2)
}
