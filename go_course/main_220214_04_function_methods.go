package main

import "fmt"

type person struct {
	name string
	age  int
}

func (aPerson person) sayGreeting() {
	fmt.Println("Good day,", aPerson.name)
}

func main() {
	jorge := person{
		name: "Jorge",
		age:  90,
	}

	jorge.sayGreeting()
}
