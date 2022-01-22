// https://www.youtube.com/watch?v=i_3O4ooSVCM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=59&ab_channel=AprendaGo
// Agrupamento de Dados
package main

import "fmt"

// arrays, slices, maps, structs

// declare an array, where its elements are of type int, and have 5 positions
var myArray [5]int

// declaration, assignment, initialization

// declaration is telling that the variable exists, and what its type
// assignment happens after declaration, we give a value to the variable
// initialization is the combined sequence of declaration and assignment

var myArray2 [6]int

func main() {
	myArray[0] = 1
	myArray[1] = 10
	fmt.Println(myArray[0], myArray[1])
	fmt.Println(myArray)
	fmt.Printf("%T\n", myArray)
	fmt.Printf("%T\n", myArray2) // this array have a different type because this is a [6]int

	// to find out the number of elements in the array
	fmt.Println(len(myArray))

	// the limitation of arrays is that the number of elements that it can use is fixed on the creation
	// we rarely use arrays in Go

}
