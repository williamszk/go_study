// https://www.youtube.com/watch?v=cjUFrS3hqgU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=70&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	myArray := [5]int{0, 1, 2, 3, 4}

	for index, value := range myArray {
		fmt.Println(index, value)
	}

	fmt.Printf("%v, %T", myArray, myArray)
}
