// https://www.youtube.com/watch?v=l2YJ-5GpGr8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=109&ab_channel=AprendaGo
package main

import "fmt"

func main() {
	x := 10

	fmt.Println("Show the variable:", x)
	fmt.Println("Show the variable's address in memory:", &x)

	// we can assign the address of x to a variable y
	y := &x
	// this is how we get the address of a variable

	fmt.Println("The value of y:", y)

	// "dereference"
	// we can use the address to find the value that is store in there
	fmt.Println("Using dereference to find out what is inside an addresss: ", *y)
	fmt.Println("Another experiment, using *&x:", *&x)

	fmt.Println()
	fmt.Printf("%T, %T", x, y) // int, *int

	// the "pointer" is the address
	// the "pointer" is a variable that stores the address of another object in memory

}
