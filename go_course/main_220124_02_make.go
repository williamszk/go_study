// https://www.youtube.com/watch?v=IMO5_ancK9w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=64&ab_channel=AprendaGo
package main

import "fmt"

func main() {
	// length
	// capacity: the array behind the slice
	mySlice := make([]int, 5, 10)

	mySlice[0], mySlice[1], mySlice[2], mySlice[3] = 1, 2, 3, 4

	for i := 0; i < 6; i++ {
		mySlice = append(mySlice, 5+i)
	}

	fmt.Println(mySlice, len(mySlice), cap(mySlice))

}
