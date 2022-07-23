// https://www.youtube.com/watch?v=clobcQqAgos&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=67&ab_channel=AprendaGo

package main

import (
	"fmt"
)

func main() {
	// maps have quick lookup
	// what would be the equivalent to the map in C?
	// when should I use map and struct
	// can we use map to represent a JSON instead of a struct

	friends := map[string]int{
		"alfred":     1234,
		"Bob Martin": 4567,
	}

	// in structs we have literals as key names
	// in maps we can choose the type: string, int

	fmt.Println(friends)

	// in maps differently from structs we use
	// square brackets to get the value from a key
	// it can be of type int, string etc
	friendExp := friends["Bob Martin"]
	fmt.Println(friendExp) // 4567

	// what happens if we access a key that does not exist in the map?
	fmt.Println(friends["Ronaldo"]) // 0

	// we need to use the comma ok idiom
	friendExp, ok := friends["Ronaldo"]
	fmt.Println("This should be not ok:", ok) // false
	fmt.Println(friendExp)
	if ok {
		fmt.Println(friendExp)
	} else {
		fmt.Println("We could not find the key:", friendExp)
	}

}

// next class:
// https://www.youtube.com/watch?v=7a6I-GnqtSE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=68&ab_channel=AprendaGo
