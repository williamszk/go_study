// https://www.youtube.com/watch?v=8Lym_4_dwOQ&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=72&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	myslice := []int{42, 43, 44, 45, 46, 47, 48, 49, 49, 50, 51}
	fmt.Println(myslice)
	myslice = append(myslice, 52)
	fmt.Println(myslice)
	myslice = append(myslice, []int{53, 54, 55}...)
	fmt.Println(myslice)

	otherslice := []int{56, 57, 58, 59, 60}
	myslice = append(myslice, otherslice...)

	fmt.Println(myslice)

}
