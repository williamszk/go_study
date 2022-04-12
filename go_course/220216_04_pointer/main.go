// https://www.youtube.com/watch?v=0slBes2RYgc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=110&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	// pointers are useful when we want to manipulate a large variable
	// in Go everything is passed as value
	// so if we want to use this variable without copying the entire object
	// we can reference it directly in memory

	x := 0

	fmt.Println(fnReceiveValue(x))
	fmt.Println(fnReceivePointer(&x))
	fmt.Println(x)

}

func fnReceiveValue(x int) int {
	// This function passes arguments by value

	x++
	return x
}

func fnReceivePointer(y *int) int {
	// This function passes arguments by reference
	*y++
	return *y
}

// next
// https://www.youtube.com/watch?v=lSAVf0RgmHc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=111&ab_channel=AprendaGo
