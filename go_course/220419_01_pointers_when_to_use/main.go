// https://www.youtube.com/watch?v=0slBes2RYgc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=110&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	x := 0

	ReceiveByValue(x)

	fmt.Println("global:", x)

	ReceivesByPointer(&x)

	fmt.Println("global:", x)

}

func ReceiveByValue(x int) {
	x++
	fmt.Println(x)
}

func ReceivesByPointer(y *int) {
	*y++
	fmt.Println(*y)
}
