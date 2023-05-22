package main

import "fmt"

// create an anonymous struct with fields that are:
// another struct, a map, a bool, a string, a slice of structs
func main() {

	var myVar = struct {
		likesInPhoto     map[string]int
		alternativeUsers []struct {
			username string
			email    string
		}
		active     bool
		sinceLogin int
	}{
		likesInPhoto: map[string]int{
			"photo01": 10,
			"photo02": 20,
			"photo03": 1,
		},
		alternativeUsers: []struct {
			username string
			email    string
		}{
			{
				username: "BobMarley",
				email:    "bob@marley.com",
			},
			{
				username: "Alice-Fergs",
				email:    "alice@crypt.so",
			},
		},
		active:     false,
		sinceLogin: 10,
	}

	fmt.Println(myVar)
	fmt.Println(myVar.active)
	fmt.Println(myVar.likesInPhoto)
	fmt.Println(myVar.alternativeUsers)
	for _, v := range myVar.alternativeUsers {
		fmt.Println(v.email, v.username)
	}
	fmt.Println("likes in photos")
	for k, v := range myVar.likesInPhoto {
		println(k, v)
	}

}
