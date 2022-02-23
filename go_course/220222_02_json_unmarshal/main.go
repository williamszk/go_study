// https://www.youtube.com/watch?v=oUsTxBwHaMM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=121&ab_channel=AprendaGo

package main

import (
	"encoding/json"
	"fmt"
)

type Person []struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	myStr := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`
	fmt.Println(myStr)

	var myJson Person

	tmpSlice := []byte(myStr)
	err := json.Unmarshal(tmpSlice, &myJson)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%v, %T\n", myJson, myJson)
}

// &myJson -> the address of the object
// "dereference" we call the object using its address (pointer)
//
