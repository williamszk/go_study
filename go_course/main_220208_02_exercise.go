package main

import "fmt"

type person struct {
	name    string
	surname string
	flavors []string
}

func main() {
	person01 := person{
		name:    "Robin",
		surname: "Hood",
		flavors: []string{"Choco", "Vanilla"},
	}

	person02 := person{
		name:    "Wonder",
		surname: "Woman",
		flavors: []string{"Jalapeno", "Belgian"},
	}

	fmt.Println(person01)
	fmt.Println(person02)
	fmt.Println(person02.flavors)

	for _, value := range person01.flavors {
		fmt.Println(value)
	}
}
