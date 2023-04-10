package main

import "fmt"

func main() {

	b := Bcon{
		A: Acon{
			A1field: "A field embedded inside Bcon",
		},
		B1field: "this is the B1field field from Bcon",
	}

	fmt.Println(b)

	fmt.Println(b.Amethod())
	fmt.Println(b.Amethod())

}

type A interface {
	Amethod() string
}

type Bcon struct {
	A
	B1field string
}

type Acon struct {
	A1field string
}

func (a Acon) Amethod() string {
	return "hi from Amethod: " + a.A1field
}
