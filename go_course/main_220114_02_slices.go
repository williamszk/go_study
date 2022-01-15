// https://www.youtube.com/watch?v=MMzTlWZ9Gjw&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=60&ab_channel=AprendaGo
// composite datatypes are built using primitive datatypes and other composite datatypes
package main

import "fmt"

func main() {
	// an example of building an array
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)
	fmt.Printf("%T\n", myArray)

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)
	fmt.Printf("%T\n", mySlice)
	// this one is an slice

	// how to append elements to an array
	mySlice2 := append(mySlice, 6)
	fmt.Println(mySlice2)
	fmt.Printf("%T\n", mySlice2)
	// we can change the size of the slice
	// we can't change the size of the array

	// we can get elements in the slice
	fmt.Println(mySlice2[3])
	// we can reassign values to elements in a slice
	mySlice2[3] = 7862
	fmt.Println(mySlice2[3])

	// the slice is based on arrays
	// when we assign a new value to an element of a slice
	// go will trash the old array and build a new one
	// this is computationaly costly then...

	// we can't do:
	// mySlice2[20] = 0

}
