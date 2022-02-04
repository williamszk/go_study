// https://www.youtube.com/watch?v=xpeExQ5C5S8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=70&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range mySlice {
		fmt.Println(index, value)
	}

	fmt.Printf("%v, %T", mySlice, mySlice)
}
