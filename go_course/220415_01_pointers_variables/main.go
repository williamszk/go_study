package main

import "fmt"

func main() {
	x := 2

	a := &x
	b := &x

	fmt.Printf("*a: %v\t*b: %v\n", *a, *b)
	fmt.Printf("Type: a: %T\tb: %T\n", a, b)
	fmt.Printf("Values a: %p\tb: %p\n", a, b)

	fmt.Println("Change dat inside b")
	*b = 99

	fmt.Printf("*a: %v\t*b: %v\n", *a, *b)

	fmt.Println("-------------- another test --------------")

	var1 := 2
	var2 := var1

	fmt.Printf("var1: %v\tvar2: %v\n", var1, var2)
	fmt.Printf("Type: var1: %T\tvar2: %T\n", var1, var2)
	fmt.Printf("Addr: var1: %p\tvar2: %p\n", &var1, &var2)

}
