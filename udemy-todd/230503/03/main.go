// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922062#overview

package main

import "fmt"

type MyType int

var x MyType

func main() {
	fmt.Printf("x=%v \t type(x)=%T\n", x, x)
	// x=0          type(x)=main.MyType
	x = 42
	fmt.Printf("x=%v \t type(x)=%T\n", x, x)
	// x=42     type(x)=main.MyType
}

// A user created type has an underlying type
// This is only valid for user create type that is just an alias
// the underlying is a primitive type like string, bool
