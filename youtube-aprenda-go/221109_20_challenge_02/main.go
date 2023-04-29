// https://www.youtube.com/watch?v=iu632z7i3MM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=140&ab_channel=AprendaGo

package main

import "fmt"

type Person struct {
	name      string
	age       int
	character string
}

func (s *Person) talkTo(o *Person) {
	fmt.Println(s.name, "is talking to", o.name)
}

type Human interface {
	talkTo(o *Person)
}

func saySomething(s Human, o *Person) {
	s.talkTo(o)
}

func main() {
	bob := Person{
		name:      "Robert Ray",
		age:       10,
		character: "Quiet",
	}

	alice := Person{
		name:      "Alice Wou",
		age:       20,
		character: "Speedy",
	}

	// this works, and it is a shortcut
	bob.talkTo(&alice)
	// and this also works, and it is the "correct" way to do it
	(&bob).talkTo(&alice)

	saySomething(&alice, &bob)

}
