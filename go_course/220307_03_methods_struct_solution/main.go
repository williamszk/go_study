package main

import "fmt"

type person struct {
	name string
	age  int
}

func (aperson *person) talk() {
	fmt.Println(aperson.name, "said hi.")
}

type humans interface {
	talk()
}

func saySomething(ahuman humans) {
	ahuman.talk()
}

func main() {

	theguy := person{
		name: "Bob Marley",
		age:  30,
	}

	theguy.talk()

	saySomething(&theguy)

}
