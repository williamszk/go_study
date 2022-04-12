// https://www.youtube.com/watch?v=G0rxcnojV_U&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=63&ab_channel=AprendaGo

package main

import "fmt"

func main() {

	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marguerita"}

	fmt.Println(sabores)

	fatia := sabores[0:]

	fmt.Println(fatia)

	// how to exclude an item from the slice?
	sabores2 := append(sabores[0:2], sabores[3:]...)

	fmt.Println(sabores2)

}

// next class
// https://www.youtube.com/watch?v=MbvABKiAABA&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=64&ab_channel=AprendaGo
