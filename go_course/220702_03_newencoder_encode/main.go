// https://www.youtube.com/watch?v=Y1Ym6Ai3uys&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=122&ab_channel=AprendaGo
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	// json.NewEncoder(os.Stdout)
	// this part is responsible to write something to the stdout
	err := json.NewEncoder(os.Stdout).Encode(u1)
	// this part of the code will save the Go object as json into os.Stdout
	// but it could be to a file also, we just need to change the os.Stdout
	// to another thing

	if err != nil {
		fmt.Println(err)
	}

}
