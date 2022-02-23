// https://www.youtube.com/watch?v=-wHOZpi697M&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=121&ab_channel=AprendaGo

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	usersJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(usersJson))

}
