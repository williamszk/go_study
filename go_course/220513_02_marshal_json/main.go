// https://www.youtube.com/watch?v=-wHOZpi697M&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=121&ab_channel=AprendaGo

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string `json:"FirstName"`
	Age   int    `json:"Age"`
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		"Moneypenny", 27,
	}

	u3 := User{"M", 54}

	users := []User{u1, u2, u3}

	fmt.Println(users)
	// [{James 32} {Moneypenny 27} {M 54}]

	outJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(outJson))
	// [{"FirstName":"James","Age":32},{"FirstName":"Moneypenny","Age":27},{"FirstName":"M","Age":54}]

}

// next:
// https://www.youtube.com/watch?v=oUsTxBwHaMM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=121&ab_channel=AprendaGo
