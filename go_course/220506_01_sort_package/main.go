// https://www.youtube.com/watch?v=b67JIGYM6Hc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=117&ab_channel=AprendaGo

package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := []string{"banana", "apple", "orange", "pear"}

	fmt.Println("Before sort: ", fruits)

	// sorting of strings
	sort.Strings(fruits)

	fmt.Println("After sort: ", fruits)

	numBought := []int{4, 3, 2, 1, 1, 2}

	fmt.Println("Before sort: ", numBought)

	sort.Ints(numBought)

	fmt.Println("After sort: ", numBought)
}
