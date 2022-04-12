// https://www.youtube.com/watch?v=2zTENBJTlD0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=91&ab_channel=AprendaGo
package main

import "fmt"

type person struct {
	name string
	age  int
}

type dentist struct {
	person
	specialization string
	wage           float64
}

type architect struct {
	person
	faculty string
	wage    float64
}

func (x dentist) sayGreeting() {
	fmt.Println("Good morning, my name is", x.name, "I like to extract teeth.")
}

func (x architect) sayGreeting() {
	fmt.Println("Good morning, my name is", x.name, "It is a good day to walk.")
}

type people interface {
	sayGreeting()
}

func humanGreeting(human people) {
	// human.sayGreeting()

	switch human.(type) {
	case dentist:
		fmt.Println("I'm a dentist specialized in ", human.(dentist).specialization)
	case architect:
		fmt.Println("I'm an architect, I studied in ", human.(architect).faculty)
	}
}

func main() {
	// an interface is a set of methods
	// an object is classified according to the set of methods that it contains

	// if types A and B have the same interface then
	// we can create a method that can receive both
	// this is called polymorphism

	archPerson := architect{
		person: person{
			name: "Howard",
			age:  76,
		},
		faculty: "Sorbonne",
		wage:    10829.00,
	}

	archPerson.sayGreeting()

	dentPerson := dentist{
		person: person{
			name: "Renata",
			age:  26,
		},
		specialization: "programming",
		wage:           999999.00,
	}

	dentPerson.sayGreeting()

	fmt.Println("====================")

	// the interface allow us to use different types as arguments to the function
	humanGreeting(archPerson)
	humanGreeting(dentPerson)
	// the ability to use the same function with different types of object
	// is called polymorphism
	// this is relevant for strongly typed languanges
	// but is not relevant for dynamically typed languanges

}

// next
// https://www.youtube.com/watch?v=pp8NKaoyefQ&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=92&ab_channel=AprendaGo
