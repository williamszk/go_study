package main

// https://www.youtube.com/watch?v=PPOBe49M8V0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=87&ab_channel=AprendaGo
import "fmt"

func add(x, y int) int {
	return x + y
}

// multiple return function
func addMany(x ...int) (int, int) {
	// inside the function x will be interpreted as a slice
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, len(x)
}

func myFunction01(mystring string, x ...int) (int, int, string) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, len(x), mystring
}

func main() {
	fmt.Println(add(10, 12))

	fmt.Println(addMany(1, 1, 1, 1, 2))

	fmt.Println(myFunction01("The phrase to be passed.", 1, 1, 1, 1, 1))
}
