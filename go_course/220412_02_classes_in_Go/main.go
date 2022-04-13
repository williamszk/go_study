package main

import "fmt"

// https://www.coderedcorp.com/blog/golang-for-python-devs/

// Fruit holds nutrition facts about a fruit.
type Fruit struct {
	Name     string
	Sugar    string
	Calories int
}

// TakeBite reduces the calories of the fruit by 10.
//
// Note that instead of `self`, we are providing the function definition with a
// reference to a `Fruit` object `f` (indicated by the asterisk `*Fruit`).
func (f *Fruit) TakeBite() {
	f.Calories -= 10
}

// note that we use *Fruit and not just Fruit
//

func main() {
	// Create a Fruit instance.
	apple := Fruit{
		Name:     "Apple",
		Sugar:    "96 grams",
		Calories: 100,
	}

	fmt.Println(apple)

	apple.TakeBite()

	fmt.Println(apple)

	apple2 := TakeBiteFunc(apple)

	fmt.Println(apple2)

}

// can we make a function that uses the struct Fruit as an argument that does the same thing as TakeBite()?
// In the case of building just a function we need to return the new object
// when we build a method the change is implicit, and it changes the instance
func TakeBiteFunc(f Fruit) Fruit {
	f.Calories -= 20

	return f
}
