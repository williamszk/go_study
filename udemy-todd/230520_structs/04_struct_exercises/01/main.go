package main

import "fmt"

// create your own type `person`
// its underlying type is struct
type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	bob := person{
		first:      "Bob",
		last:       "Ferguson",
		favFlavors: []string{"Chocolate", "Strawberry"},
	}

	alice := person{
		first:      "Alice",
		last:       "Gee",
		favFlavors: []string{"Ice", "Fire"},
	}

	fmt.Printf("%s %s's favorite ice cream flavors are: ", bob.first, bob.last)
	for _, flavor := range bob.favFlavors {
		fmt.Printf("%s, ", flavor)
	}
	fmt.Println()

	fmt.Printf("%s %s's favorite ice cream flavors are: ", alice.first, alice.last)
	for _, flavor := range alice.favFlavors {
		fmt.Printf("%s, ", flavor)
	}
	fmt.Println()

}
