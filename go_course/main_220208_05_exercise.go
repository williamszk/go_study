package main

import "fmt"

func main() {
	myStruct := struct {
		mychars map[string]string
		points  []int
	}{
		mychars: map[string]string{
			"color": "blue",
			"speed": "102",
		},
		points: []int{10, 12, 15, 18},
	}

	fmt.Println(myStruct)
}

// next
// https://www.youtube.com/watch?v=PPOBe49M8V0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=87&ab_channel=AprendaGo
