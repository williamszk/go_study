// https://www.youtube.com/watch?v=l2YJ-5GpGr8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=109&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println("The address of the variable:", &x)

	// define a pointer
	y := &x
	// here y is a pointer, a pointer is a variable in itself
	// and the pointer have an address

	fmt.Println("Print the value of the pointer (which is a variable):", y)

	// we can also make a dereference
	// which means that with a pointer we can access the value that is stored in the address
	// the address that the pointer is pointing to

	fmt.Println("An example of dereferencing:", *y)

	// another way that we could do the same thing is:
	fmt.Println("Dereferencing the address of an address:", *&x)

	//
	fmt.Printf("What is the type of a pointer?\t %v\t%T\n", y, y) // 0xc0000aa058   *int
	// it is a *int
	// In C it is int*

	// An example
	a := 0
	b := &a
	*b++
	fmt.Println("What is the expected result if we print a?", a)
}

// next video:
// https://www.youtube.com/watch?v=l2YJ-5GpGr8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=110&ab_channel=AprendaGo
