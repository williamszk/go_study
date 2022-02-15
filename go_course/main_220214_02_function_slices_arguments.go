package main

import "fmt"

// https://www.youtube.com/watch?v=pV5OEqIRPh4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=88&ab_channel=AprendaGo

func add(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func main() {
	// how to pass a slice as argument to a function?
	myslice := []int{1, 2, 3, 4}
	fmt.Println(add(myslice...))
	// it is like a pointer, but more like *args in python
}
