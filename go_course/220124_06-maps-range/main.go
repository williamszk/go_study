// https://www.youtube.com/watch?v=7a6I-GnqtSE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=68&ab_channel=AprendaGo
package main

import "fmt"

func main() {
	myMap := map[int]string{
		123: "muito legal",
		98:  "menos legal",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	}

	fmt.Println(myMap)

	// using range to read key value in a map
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// how to delete elements from a map?
	delete(myMap, 123)

	fmt.Println(myMap)
}

// next
// https://www.youtube.com/watch?v=cjUFrS3hqgU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=69&ab_channel=AprendaGo
