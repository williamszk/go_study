// https://www.youtube.com/watch?v=f4lvrtXsGIM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=75&ab_channel=AprendaGo

package main

import "fmt"

func main() {

	tableUser := [][]string{
		{"Name", "Surname", "Hobby"},
		{"Jane", "Mary", "Jump"},
		{"Peter", "Parker", "Walk"},
		{"Harry", "Potter", "Clean"},
	}

	fmt.Println(tableUser)

	for index, line := range tableUser {
		fmt.Printf("%v", index)
		for _, cell := range line {
			fmt.Printf("\t%v\t", cell)
		}
		fmt.Printf("\n")
	}

}
