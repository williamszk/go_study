// https://www.youtube.com/watch?v=3hidzcEZ0KE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=124&ab_channel=AprendaGo

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type orderByFirstName []user

func (users orderByFirstName) Len() int {
	return len(users)
}

func (users orderByFirstName) Less(i, j int) bool {
	return users[i].First < users[j].First
}

func (users orderByFirstName) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

type orderByLastName []user

func (users orderByLastName) Len() int {
	return len(users)
}

func (users orderByLastName) Less(i, j int) bool {
	return users[i].Last < users[j].Last
}

func (users orderByLastName) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

type orderByAge []user

func (users orderByAge) Len() int {
	return len(users)
}

func (users orderByAge) Less(i, j int) bool {
	return users[i].Age < users[j].Age
}

func (users orderByAge) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(orderByFirstName(users))

	// fmt.Println("Sort by First Name:", users)

	sort.Sort(orderByLastName(users))

	// fmt.Println("Sort by Last Name:", users)

	sort.Sort(orderByAge(users))

	// fmt.Println("Sort by Age:", users)

	// order the Sayings

	for i, v := range users {
		fmt.Println(i, "------------------------")
		fmt.Println("First Name:", v.First)
		fmt.Println("Last Name:", v.Last)
		fmt.Println("Sayings:")
		for j, saying := range v.Sayings {
			fmt.Println("\t", j, saying)
		}

	}

}

// next
// https://www.youtube.com/watch?v=unofca9ooS4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=125&ab_channel=AprendaGo
// concurrency
