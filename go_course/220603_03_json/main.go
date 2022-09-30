// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922290#overview

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// what happens if we define struct inside the main func?
	// we can use it in the same way;
	// this struct is used to group data related to color palettes...
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	// the input type of Marshal is an interface{}
	// idk what that means
	b, err := json.Marshal(group)
	// the retun of Marshal is a slice of bytes
	// the Marshal function will transform a type struct of some format into a string
	// this string is of json format
	// we can save this string into a file or send
	// as a response to an API

	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(b)
	// the difference between using os.Stdout.Write and fmt.Println is that
	// the former will print the slice of byte in string format
	// I think that fmt.Println uses os.Stdout.Write behind the scenes
	fmt.Println("")
	fmt.Println(b)
	fmt.Println(string(b))

	fmt.Println("==============================================================")

	// What if we want to send a list of names to the JSON?
	// or a list of ColorGroup?

	myList := []string{"Clean Code", "Clean Coder", "Clean Architecture", "Software Craftsmanship"}
	// the input argument of the Marshal function
	// is an interface{} this could be a slice of user defined struct
	// or it can be just a struct
	// or it can be a slice of string, or a slice of floats
	// it can be many things
	b, err = json.Marshal(myList)

	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(b)
	// this worked
	fmt.Println("")
	fmt.Println("==============================================================")

	AListOfColors := []ColorGroup{
		{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		},
		{
			ID:     2,
			Name:   "Blues",
			Colors: []string{"Blue Navy", "Blue", "Dark Blue", "Turquoise"},
		},
	}
	fmt.Println("In the usual format")
	fmt.Println(AListOfColors)

	b, err = json.Marshal(AListOfColors)

	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(b)
	fmt.Println("")
	fmt.Println("==============================================================")
}
