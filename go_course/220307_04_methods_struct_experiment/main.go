package main

import (
	"fmt"
)

type dentist struct {
	name           string
	age            int
	specialization string
}

type architect struct {
	name     string
	age      int
	building string
}

func (aDentist dentist) talk() {
	fmt.Println(aDentist.name, "said hi. She is specialized in", aDentist.specialization)
}

func (anArchitect architect) talk() {
	fmt.Println(anArchitect.name, "said hi. She builds", anArchitect.building)
}

type people interface {
	talk()
}

func saySomething(aPerson people) {
	aPerson.talk()
}

func main() {

	dent1 := dentist{
		name:           "Mary",
		age:            60,
		specialization: "Orto",
	}

	dent1.talk()

	arch1 := architect{
		name:     "Jane",
		age:      70,
		building: "Houses",
	}

	arch1.talk()

	fmt.Println("Using function saySomething():")
	saySomething(dent1)
	saySomething(arch1)

}

// next:
// https://www.youtube.com/watch?v=cCWvFijhObU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=141&ab_channel=AprendaGo
