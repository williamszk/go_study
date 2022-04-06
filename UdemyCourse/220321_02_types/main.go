// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922046#overview
// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922050#overview
package main

import "fmt"

type hotdog int // the underlying type of hotdog is an int
var b hotdog

func main() {
	a := 42
	fmt.Printf("%v\t%T\n", a, a)

	b = 43
	fmt.Printf("%v\t%T\n", b, b)

	// a = b // this is not allowed
	// b = a // we cannot do this too

	a = int(b) // this is allowed
	// this is called type conversion and not type casting as in other programming languages
	fmt.Printf("a = %v\t%T\n", a, a)
	fmt.Printf("b = %v\t%T\n", b, b)
}
