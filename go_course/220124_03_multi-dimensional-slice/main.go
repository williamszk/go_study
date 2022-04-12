// https://www.youtube.com/watch?v=o3yoYGWqrDE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=65&ab_channel=AprendaGo
package main

import "fmt"

func main() {

	sliceofslices := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(sliceofslices)
	fmt.Println(sliceofslices[0][1]) // should be 2
}
