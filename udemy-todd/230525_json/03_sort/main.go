package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{4, 7, 3, 42, 99, 16}
	xs := []string{"James", "Q", "Moneypenny", "Dr."}

	fmt.Println("before", xi)
	sort.Ints(xi)
	fmt.Println("after", xi)

	fmt.Println("before", xs)
	sort.Strings(xs)
	fmt.Println("after", xs)
}
