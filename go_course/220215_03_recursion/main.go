// https://www.youtube.com/watch?v=1-pop5h5RAs&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=97&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	fmt.Println("Test factorial of 4 is", factorial(4))
	fmt.Println("Test factorial of 5 is", factorial(5))
	fmt.Println("Test factorial of 5 is", factorial(6))

	fmt.Println("Test loops of 4 is", loops(4))
	fmt.Println("Test loops of 5 is", loops(5))
	fmt.Println("Test loops of 5 is", loops(6))

}

func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return factorial(x-1) * x
	}
}

/*This is a factorial with while loop*/
func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
