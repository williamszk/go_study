package main

import "fmt"

type MyType int

var y int

func main() {

	fmt.Printf("%v \t %T\n", y, y)
	y = 90
	fmt.Printf("%v \t %T\n", y, y)

	y2 := MyType(y)
	fmt.Printf("%v \t %T\n", y2, y2)

	x := MyType(10)
	fmt.Printf("%v \t %T\n", x, x)

	y = int(x)
	fmt.Printf("%v \t %T\n", y, y)
}
