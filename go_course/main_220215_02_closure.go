// https://www.youtube.com/watch?v=mOM0qTB5ppU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=96&ab_channel=AprendaGo

package main

import "fmt"

func main() {

	a := i()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	// x exists and keeps being inside the scope of function a
	// this is not good because the same function is returning different things
	// this is not a pure function

	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0 // scoped variable, this is a captured variable inside the closure
	return func() int {
		x++
		return x
	}
}
