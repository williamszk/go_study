// https://www.youtube.com/watch?v=l2YJ-5GpGr8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=109&ab_channel=AprendaGo
package main

import "fmt"

func main() {

	// an experiment with pointers
	x := 0

	y := &x

	fmt.Println(x, y)

	// increment the value of x directly through the pointer
	*y++

	fmt.Println(x, y)

}
