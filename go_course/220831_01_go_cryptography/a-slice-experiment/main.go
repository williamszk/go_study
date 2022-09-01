package main

import "fmt"

func main() {

	var MyArray [5]float32

	MySlice := []int{1, 2, 3, 4}

	fmt.Println(MySlice)
	fmt.Println(MyArray)

	// how to access the capacity of the slice
	// how to know the length of the slice

	fmt.Println("Capacity of slice", cap(MySlice))
	fmt.Println("Length of slice", len(MySlice))

	fmt.Printf("type of MyArray: %T\n", MyArray)    // [5]float32 this is an array
	fmt.Printf("content of MyArray: %v\n", MyArray) // [0 0 0 0 0]
	MySli01 := MyArray[:]
	fmt.Printf("type of MySli01: %T\n", MySli01)    // []float32 this is an slice
	fmt.Printf("content of MySli01: %v\n", MySli01) // [0 0 0 0 0]

	MySli02 := MyArray[:]
	fmt.Printf("type of MySli02: %T\n", MySli02)    // []float32 this is an slice
	fmt.Printf("content of MySli02: %v\n", MySli02) // [0 0 0 0 0]

	fmt.Println("Change the MySli02")
	MySli02[0] = 999.00
	fmt.Printf("content of MySli02: %v\n", MySli02) //  [999 0 0 0 0]
	fmt.Printf("content of MySli01: %v\n", MySli01) //  [999 0 0 0 0]

	// note thtat the slice is a reference to an underlying array
	// if we create two slices sli01 and sli02 from the same array
	// if we change one slice we change the other

}
