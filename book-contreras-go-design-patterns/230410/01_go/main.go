// An experiment with closures.
// A closure is an anonymous function. But it carries some information.
// That is set during its calling not its definition.
// In the case below addN is a closure.
// We can branch this addN into multiple specific implementations of functions
// when we change the argument.
// It is called a closure because if encloses some local information.
package main

import "fmt"

func main() {

	addN := func(m int) func(int) int {
		inner := func(n int) int {
			return m + n
		}
		return inner
	}

	addFive := addN(5)
	res1 := addFive(2)

	fmt.Println("res1:", res1) // 7

	res2 := addFive(5)
	fmt.Println("res2:", res2) // 10

	addN2 := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}

	addTen := addN2(10)
	res3 := addTen(2)
	fmt.Println("res3:", res3) // 12

}
