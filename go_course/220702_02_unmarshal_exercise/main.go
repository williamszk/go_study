// https://www.youtube.com/watch?v=oUsTxBwHaMM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=122&ab_channel=AprendaGo

package main

import (
	"encoding/json"
	"fmt"
)

// look at json to go
// https://mholt.github.io/json-to-go/
type LordOfTheRingsPerson struct {
	Name        string `json:"Name"`
	Surname     string `json:"Surname"`
	Age         int    `json:"Age"`
	Profession  string `json:"Profession"`
	Bankbalance int    `json:"Bankbalance"`
}

type Movie007Person []struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	// we have a string that is in json format
	// then we can transform it into a Go native data structure
	myjsondata1 := `{"Name":"Faramir","Surname":"Denethorson","Age":34,"Profession":"Knight of Gondor","Bankbalance":100000}`

	myjsondata2 := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`

	var Faramir LordOfTheRingsPerson

	err := json.Unmarshal([]byte(myjsondata1), &Faramir)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v, %T\n", Faramir, Faramir)

	var JamesBond Movie007Person

	err = json.Unmarshal([]byte(myjsondata2), &JamesBond)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v, %T\n", JamesBond, JamesBond)
}
