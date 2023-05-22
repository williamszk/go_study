package main

import "fmt"

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

	// fmt.Printf("%s %s's favorite ice cream flavors are: ", bob.first, bob.last)
	// for _, flavor := range bob.favFlavors {
	// 	fmt.Printf("%s, ", flavor)
	// }
	// fmt.Println()

	// fmt.Printf("%s %s's favorite ice cream flavors are: ", alice.first, alice.last)
	// for _, flavor := range alice.favFlavors {
	// 	fmt.Printf("%s, ", flavor)
	// }
	// fmt.Println()

	var peopleAndFavFlavors = make(map[string]person)

	peopleAndFavFlavors[bob.last] = bob
	peopleAndFavFlavors[alice.last] = alice

	for _, v := range peopleAndFavFlavors {
		fmt.Printf("%s %s's favorite ice cream flavors are: ", v.first, v.last)
		for _, flavor := range v.favFlavors {
			fmt.Printf("%s, ", flavor)
		}
		fmt.Println()
	}

}
