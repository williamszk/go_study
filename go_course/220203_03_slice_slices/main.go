// https://www.youtube.com/watch?v=Wzc5VOjh-XQ&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=71&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("First to third item included:", mySlice[0:3])
	fmt.Println("From fifth to the last item:", mySlice[4:])
	fmt.Println("Second to seventh item included:", mySlice[1:6])
	fmt.Println("Third to before the last one included:", mySlice[2:len(mySlice)-1])

}
