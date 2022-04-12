// https://www.youtube.com/watch?v=gA1cTnWjPaU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=73&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	myslice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(myslice)

	myslice = append(myslice[:3], myslice[6:]...)

	fmt.Println(myslice)

}
