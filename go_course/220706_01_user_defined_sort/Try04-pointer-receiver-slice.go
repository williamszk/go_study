// Here we experiment with building a compost object
// a slice, and test if there is implicit change of the object
// with method or usual functions.
// That is, if the functions will pass the argument by value or
// they will pass by reference.
// With structs or atomic objects (int, string), we know that
// we need to use pointer receiver

// Take a look at this notes:
// https://go.dev/tour/methods/4

package main

import "fmt"

func SwapOrder(m []int, i, j int) {
	m[i], m[j] = m[j], m[i]
}

func Try04() {
	myList := []int{1, 2, 3, 4}

	fmt.Println("Before swap: ", myList)

	SwapOrder(myList, 0, 1)

	fmt.Println("After swap: ", myList)

	// We can see that we in case of slices the change is permanent
	// we don't need to pass the argument using pointer receiver
	// Before swap:  [1 2 3 4]
	// After swap:  [2 1 3 4]
}

// From Try03 we saw that structs are passed as value by defaul
// but in here we can see that slices are not

// What happens then if we have a slice inside a struct?
// In this case I think that the slice will have a permanent "change"
// given that we are actually changing the backing array of the slice

// https://stackoverflow.com/a/39993797/8782077
// Everything in Go is passed by value, slices too. But a slice value is a header, describing a contiguous section of a backing
// array, and a slice value only contains a pointer to the array where the elements are actually stored.
// The slice value does not include its elements (unlike arrays).
// So when you pass a slice to a function, a copy will be made from this header, including the pointer, which will point to the
// same backing array. Modifying the elements of the slice implies modifying the elements of the backing array, and so
// all slices which share the same backing array will "observe" the change.
