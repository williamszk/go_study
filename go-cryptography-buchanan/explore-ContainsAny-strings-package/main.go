// Explore the functionality of function ContainsAny from the the package
// strings

package main

import (
	"fmt"
	"strings"
)

func main() {

	// func isHexString(s string) bool {
	// 	return strings.ContainsAny(s, "abc")
	// }

	MyStr := "Hello World"

	contains := strings.ContainsAny(MyStr, "H") // true
	fmt.Println("contains: ", contains)

	contains = strings.ContainsAny(MyStr, "h") // false
	fmt.Println("contains: ", contains)

	contains = strings.ContainsAny(MyStr, "a") // false
	fmt.Println("contains: ", contains)

	MyStr = "def"
	contains = strings.ContainsAny(MyStr, "abc") // false
	fmt.Println("contains: ", contains)

	MyStr = "0x7be29a"
	contains = strings.ContainsAny(MyStr, "abc") // true
	fmt.Println("contains: ", contains)

	MyStr = "7ef"
	contains = strings.ContainsAny(MyStr, "abc") // false
	fmt.Println("contains: ", contains)
}
