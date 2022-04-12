// https://www.youtube.com/watch?v=dRNNC7VpztE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=66&ab_channel=AprendaGo
package main

import "fmt"

func main() {
	primeiroslice := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiroslice)
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)
	// note that the primeroslice has changed...
	// this is not desirable
}

// https://www.youtube.com/watch?v=clobcQqAgos&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=67&ab_channel=AprendaGo
// next class
