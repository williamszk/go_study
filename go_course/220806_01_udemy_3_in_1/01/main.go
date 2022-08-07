package main

import "fmt"

func main() {

	// this is an array
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	b := [10]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a}

	// this is a slice
	c := []float32{1, 2, 3, 4}

	c = append(c, []float32{5, 6, 7}...)
	// the ... elipsis is the spread operator

	c = append(c, 0)
	// c = append(c, "a")

	// we say that append is a variadic function because it
	// accepts a variable number of parameters
	c = append(c, 1, 2, 3)
	c = append(c, 0.0, 0.0, 0.0, 0.0, 0xff, 0x7f)

	// initiallize an empty slice of uint64
	c2 := []uint64{}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%v, %T\n", c2, c2)

	// experiment with underling array of a slice
	myslice := make([]int, 5) // 5 is the capacity and the size of the slice
	fmt.Println(myslice)      // [0 0 0 0 0]
	for i := 1; i <= 5; i++ {
		myslice[i-1] = i
	}
	fmt.Println(myslice)
	myslice1 := myslice[1:3] // [1 2 3 4 5]
	fmt.Println(myslice1)    // [2 3]
	fmt.Println(cap(myslice1))

	// change the underlying value of the array and see if the sub slice is affected
	myslice[1] = 99
	fmt.Println(myslice1) // [99 3]
	// the sub slice is affected

	myslice2 := make([]int, 2)
	copy(myslice2, myslice[1:3])
	fmt.Println((myslice2)) // [99 3]
	// myslice2 shouldn't change

	myslice[1] = 2
	fmt.Println((myslice2)) // [99 3] <- this did not change because the underlying array didn't change
	fmt.Println((myslice1)) // [2 3] <- this changed because the underlying array changed

	fmt.Printf("&myslice =  %p\n", &myslice)
	fmt.Printf("&myslice1 = %p\n", &myslice1)
	fmt.Printf("&myslice2 = %p\n", &myslice2)

}
