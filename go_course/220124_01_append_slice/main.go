// https://www.youtube.com/watch?v=MbvABKiAABA&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=64&ab_channel=AprendaGo
package main

import "fmt"

func main() {

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}
	// how to join two slices
	umaslice = append(umaslice, outraslice...)
	// this is like the spread from javascript
	// this can be called unfurl in Go
	// and can be called enumeration of the slice

	fmt.Println(umaslice)
}
