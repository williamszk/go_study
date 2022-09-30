package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type OrderByLastname []User

func (s OrderByLastname) Len() int {
	return len(s)
}

func (s OrderByLastname) Less(i, j int) bool {
	return s[i].Last < s[j].Last
}

// I have the impression that methods make inplace transformations of the
// the object
//
func (s OrderByLastname) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Try01() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken",
			"Youth",
			"In",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"it is",
			"Would",
			"I would",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh",
			"Dear",
			"Can",
		},
	}

	users := []User{u1, u2, u3}

	// Order users by Last name and age

	fmt.Println("Before sort.Strings()", u1.Sayings)
	fmt.Println(u1.Sayings)

	sort.Strings(u1.Sayings)
	// the sort.Strings happens inplace
	fmt.Println("After sort.Strings()", u1.Sayings)
	// sort.Sort()

	// Are methods always applied inplace?
	UsersByLast := OrderByLastname(users)

	fmt.Println("Before Swap: ", users)
	UsersByLast.Swap(0, 1)
	fmt.Println("After Swap: ", users)

	sort.Sort(UsersByLast)
	fmt.Println("After sort.Sort: ", users)

}
