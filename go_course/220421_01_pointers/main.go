// https://www.youtube.com/watch?v=lSAVf0RgmHc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=112&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=XVd-y0t5fno&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=112&ab_channel=AprendaGo
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	a := 10

	fmt.Printf("The address is: %p\n", &a)

	person01 := person{
		name: "Bob",
		age:  40,
	}

	fmt.Println(person01)

	changeNameToPeter(&person01)

	fmt.Println(person01)
}

func changeNameToPeter(p *person) {
	(*p).name = "Peter"
}

// Next:
// https://www.youtube.com/watch?v=jnnIgvV0_yA&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=113&ab_channel=AprendaGo
