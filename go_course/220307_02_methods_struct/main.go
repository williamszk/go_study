package main

import "fmt"

type person struct {
	name string
	age  int
}

func (pntPerson *person) talk() {
	fmt.Println("The person", pntPerson.name, "talked.")
}

type humans interface {
	talk()
}

func saySomething(aPerson humans) {
	aPerson.talk()
}

func main() {

	ruppert := person{
		name: "Ruppert",
		age:  10,
	}

	pntPerson := &ruppert

	fmt.Printf("ruppert: %v, %T\n", ruppert, ruppert)
	fmt.Printf("pntPerson: %v, %T\n", pntPerson, pntPerson)

	ruppert.talk()

	saySomething(&ruppert)

}
