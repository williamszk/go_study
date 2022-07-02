// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922050?start=0#overview

package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b // this will give an error
	a = int(b)
	// int(b) is of type int

	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
