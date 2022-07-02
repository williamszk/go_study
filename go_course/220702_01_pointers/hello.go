// https://www.youtube.com/watch?v=-FiBp1OeZF0&ab_channel=FullCycle

package main

import "fmt"

func main() {

	a := 10

	fmt.Printf("Address in memory of a: 0x%X\n", &a)
	// we have two options to print the memory address of a variable
	fmt.Printf("Address in memory of a: %v\n", &a)

	// creating a pointer
	var ptrA *int = &a

	fmt.Printf("A pointer to A: %v, with address: 0x%X\n", ptrA, &ptrA)

	// example of dereferencing
	fmt.Printf("Dereferencing a pointer *ptrA = %v\n", *ptrA)

	// changing the value of a using the pointer
	*ptrA = 50
	fmt.Printf("The value of a after change using pointer: %v\n", a)

	val1 := SetTo200(ptrA)
	// SetTo200 will make an implicit change in a, using its pointer
	fmt.Printf("Value of a after func SetTo200: %v\n", a)
	fmt.Println("val1 = ", val1)

	SetTo500(a)
	fmt.Printf("Value of a after func SetTo500: %v\n", a)
	// this will be 200 still, because in Go arguments are passed
	// as values into functions

	SetTo600(a)
	// when we try to print the address of "a" inside the function
	// we see that it is not the same address:
	// original "a" address:				0xC0000BA000
	// "a" inside the function address:

	MyCar := Car{
		Name: "Volkswagen",
	}

	fmt.Printf("Memory address of MyCar: \t0x%X\n", &MyCar.Name)
	MyCar.move()
	fmt.Printf("MyCar Name after change: %v\n", MyCar.Name)
	// this will still return Volkswagen

	MyCar.moveImplict()
	fmt.Printf("(After moveImplict()) MyCar Name after change: %v\n", MyCar.Name)
}

type Car struct {
	Name string
}

func (c Car) move() {
	fmt.Printf("Memory address of \"c\": \t\t0x%X\n", &c.Name)
	c.Name = "BMW"
	fmt.Println(c.Name, " is moving")
}

func (c *Car) moveImplict() {
	// When we use *Car the type of pointer to car, we specify
	// that we are passing the argument by reference
	fmt.Printf("(Inside moveImplict()) Memory address of \"c\": \t\t0x%X\n", &c.Name)
	c.Name = "BMW"
	fmt.Println(c.Name, " is moving")
}

func SetTo200(p *int) int {
	// a is of type *int, it a pointer to int
	*p = 200
	// this will make a implicit change in a
	return *p
}

// is there a way to make implicit alterations in "a" without using pointers?

func SetTo500(a int) {
	// this will not work to change implicitly the value of "a"
	a = 500
}

func SetTo600(a int) {
	fmt.Printf("Address of a passed as argument to SetTo600: 0x%X\n", &a)
}
