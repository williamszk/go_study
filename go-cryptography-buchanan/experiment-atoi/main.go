// The objective of this program is to experiment with the function
// Atoi, strconv
// Itoa

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// we will use Atoi or Itoa when we want to conver integers to strings
	// and the opposite

	// why can't we use a simple int or string to transform them into int and string?

	MyVar01 := 42

	fmt.Printf("MyVar01: %v (%T)\n", MyVar01, MyVar01) // MyVar01: 42 (int)

	// transform int into string
	MyVar02 := string(MyVar01)                         // <-- this even gives a warning
	fmt.Printf("MyVar02: %v (%T)\n", MyVar02, MyVar02) // MyVar02: * (string)
	// this does not seem right

	MyVar03 := strconv.Itoa(MyVar01)
	fmt.Printf("MyVar03: %v (%T)\n", MyVar03, MyVar03) // MyVar03: 42 (string)
	// this seems to be correct

	// an alternative
	MyVar04 := fmt.Sprintf("%v", MyVar01)
	fmt.Printf("MyVar04: %v (%T)\n", MyVar04, MyVar04) // MyVar04: 42 (string)
	// this gives the same result as strconv.Atoi

	// henve we have two alternatives to transform an int into a string
	// fmt.Sprintf is more general into transforming anything into string though
	// and strconv.Itoa is more specifif to transform int into string

	MyVarStr := "109918"
	fmt.Printf("MyVarStr: %v (%T)\n", MyVarStr, MyVarStr) // MyVarStr: 109918 (string)

	// this function Atoi will transform the string into int
	MyVar10, _ := strconv.Atoi(MyVarStr)
	fmt.Printf("MyVar10: %v (%T)\n", MyVar10, MyVar10) // MyVar10: 109918 (int)

}
