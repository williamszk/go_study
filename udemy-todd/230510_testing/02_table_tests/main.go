package main

import "fmt"

func main() {
	fmt.Println("1 + 2 + 3 =", mySum(1, 2, 3))
	fmt.Println("4 + 7 =", mySum(4+7))
	fmt.Println("5 + 9 =", mySum(5+9))
}

// variadic parameter: n number of parameters
func mySum(xi ...int) int {
	// xi is a slice of ints
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}
