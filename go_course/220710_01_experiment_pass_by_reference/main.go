package main

import "fmt"

func main() {

	doggy := Animal{
		Name: "Doggy",
		Size: 10,
	}

	fmt.Println(doggy)
	fmt.Printf("Calling from outside: \t\t\t\t\t%p\n", &doggy)

	HandleAnimal(doggy)

	doggy.HandleFromInside()

	doggy.HandleFromInsidePtrReceiver()

	// Calling from outside:                                   0xc00000c030
	// Calling from inside HandleAnimal():                     0xc00000c060
	// Calling from inside HandleFromInside():                 0xc00000c078
	// Calling from inside HandleFromInsidePtrReceiver() ptr:  0xc00000e030
	// Calling from inside HandleFromInsidePtrReceiver():      0xc00000c030

}

func HandleAnimal(a Animal) {

	fmt.Printf("Calling from inside HandleAnimal(): \t\t\t%p\n", &a)

}
