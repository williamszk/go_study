// https://www.youtube.com/watch?v=pp8NKaoyefQ&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=92&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=j9C66R4BMWM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=93&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=9Oxmya_A-Sc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=94&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=u8qBzOAhbRM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=95&ab_channel=AprendaGo

package main

import "fmt"

func main() {

	x := 10

	// this is an anonymous function
	func(x int) {
		fmt.Println(x)
	}(x)

	// this is a function as an expression

	y := func(x int) {
		fmt.Println("I am a function that is an expression.", x)
	}

	y(x) // call the function

	// how to return a function
	f1 := myHigherFunc(2)
	fmt.Println("Calling the function that is output from the higher order function:", f1(10))

	myResponse := onlyEven(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	fmt.Println(myResponse)

}

func myHigherFunc(x int) func(int) int {
	// I think that this can be called a closure
	return func(i int) int {
		return x * i
	}
}

// callback is doen when passing a function as an argument

/*
This is like a reduce function
*/
func sum(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

/*
This acts like a filter and apply a reducer.
This is an example of how we can pass a function as an argument to another function.
argfunc is a callback function.
*/
func onlyEven(argfunc func(args ...int) int, argsOut ...int) int {
	var slice []int
	for _, v := range argsOut {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := argfunc(slice...)
	return total
}

func onlyOdd(argfunc func(args ...int) int, argsOut ...int) int {
	var aslice []int
	for _, v := range argsOut {
		if v%2 != 0 {
			aslice = append(aslice, v)
		}
	}

	total := argfunc(aslice...)
	return total
}
