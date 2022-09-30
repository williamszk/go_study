// https://www.youtube.com/watch?v=fZZclCKnr7k&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=123&ab_channel=AprendaGo
package main

import (
	"fmt"
	"sort"
)

func main() {

	listnums := []int{5, 8, 2, 43, 987, 14, 12, 21, 1, 4, 4, 2, 93, 13}
	listwords := []string{"random", "raibow", "delights", "in", "torpedo", "summers", "under", "gallantry"}
	listfloats := []float64{1.2, 44.2, 55.6, 1.8, 987.3, 32.1}

	fmt.Printf("Before sort %v\n", listnums)
	sort.Ints(listnums)
	fmt.Printf("After sort %v\n", listnums)

	fmt.Printf("Before sort %v\n", listwords)
	sort.Strings(listwords)
	fmt.Printf("After sort %v\n", listwords)

	fmt.Printf("Before sort %v\n", listfloats)
	sort.Float64s(listfloats)
	fmt.Printf("After sort %v\n", listfloats)

}
