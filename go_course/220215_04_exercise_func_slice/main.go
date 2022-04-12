package main

import "fmt"

func main() {

	myslice := []int{1, 2, 3, 4, 5}
	fmt.Println(sumMany(myslice...))

	fmt.Println(sumSlice(myslice))

}

/* We can pass many arguments into a function*/
func sumMany(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

/* We can a whole slice into a function */
func sumSlice(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// next
// https://www.youtube.com/watch?v=3pbmLfcgN2s&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=102&ab_channel=AprendaGo
