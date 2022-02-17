/*
How to change fields of an struct using pointers inside functions
*/
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeMe(adrPerson *Person) {
	(*adrPerson).age = 999
	(*adrPerson).name = "Changed Name"
}

func main() {

	mike := Person{
		name: "Mike",
		age:  10,
	}

	fmt.Println(mike.name, mike.age)

	changeMe(&mike)

	fmt.Println(mike.name, mike.age)

}
