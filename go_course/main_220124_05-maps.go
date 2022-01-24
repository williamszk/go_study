// https://www.youtube.com/watch?v=clobcQqAgos&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=68&ab_channel=AprendaGo
package main

import "fmt"

func main() {
	amigos := map[string]int{
		"alfredo": 5555555,
		"joana":   11111111,
	}

	// how to access elements from the map
	fmt.Println(amigos["joana"])
	fmt.Println(amigos)

	// add new amigos
	amigos["gopher"] = 444444

	// what happens when we try to access a value from a key that does not exist?
	fmt.Println(amigos["romario"])
	// it shows 0

	thereIs, isOk := amigos["ghost"]
	fmt.Println(thereIs, isOk)

	// this is the comma ok idiom
	if thereIs, isOk = amigos["ghost"]; !isOk {
		fmt.Println("Thre is not Ghost!")
	} else {
		fmt.Println("There is Ghost: ", thereIs)
	}

}
