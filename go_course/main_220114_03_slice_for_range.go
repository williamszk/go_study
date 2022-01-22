// https://www.youtube.com/watch?v=l6O8HWaoy_w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=61&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	mySlice := []string{"banana", "maça", "jaca"}

	for index, value := range mySlice {
		fmt.Println("In index", index, "we got", value)
	}

	// if we give a new value this way bad things will happen
	// mySlice[3] = "melancia"
	mySlice = append(mySlice, "melancia")
	for index, value := range mySlice {
		fmt.Println("In index", index, "we got", value)
	}
	// it is not good to expand a slice by using element assignment, instead we use append()

	mySlice2 := []int{23, 24, 54}
	totalValue := 0
	for _, value := range mySlice2 {
		totalValue += value
	}

	fmt.Println("The total value is:", totalValue)
}

// next class
// https://www.youtube.com/watch?v=G0rxcnojV_U&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=62&ab_channel=AprendaGo



