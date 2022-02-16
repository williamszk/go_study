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

	myMap := map[string]person{
		person01.surname: person01,
		person02.surname: person02,
	}

	for _, vPerson := range myMap {
		for _, vItem := range vPerson.flavors {
			fmt.Println(vItem)
		}
	}
}

// The map is closer to dictionries and objects, instead of an anonymos struct
