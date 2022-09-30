// https://www.youtube.com/watch?v=Y1Ym6Ai3uys&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=122&ab_channel=AprendaGo

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// https://pkg.go.dev/encoding/json
	// https://pkg.go.dev/encoding/json#NewEncoder
	// https://pkg.go.dev/io#Writer
	// func NewEncoder(w io.Writer) *Encoder
	// https://golang.hotexamples.com/examples/github.com.ugorji.go.codec/-/NewEncoder/golang-newencoder-function-examples.html

	// MyEncoder := json.NewEncoder().Encode()

	// fmt.Println(MyEncoder)

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

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// func NewEncoder(w io.Writer) *Encoder
	// w io.Writer : e.g. os.Stdout, this will send the values to the terminal
	// *Encoder : this is a struct
	// func (*Encoder) Encode
	// func (enc *Encoder) Encode(v any) error
	// the function is called Encode
	// this is a method of the *Encoder type
	// the function Encode will return error

	EncoderPtr := json.NewEncoder(os.Stdout)
	fmt.Printf("os.Stdout = %v, %T\n", os.Stdout, os.Stdout)
	// os.Stdout = &{0xc00007e280}, *os.File
	fmt.Printf("EncoderPtr = %v, %T\n", EncoderPtr, EncoderPtr)
	// EncoderPtr = &{0xc000006018 <nil> true <nil>  }, *json.Encoder

	err := EncoderPtr.Encode(users)
	// I think the above line will already print the json to the console....
	if err != nil {
		fmt.Println(err)
	}
	// [
	//	{
	//	"First":"James",
	//	"Last":"Bond",
	// 	"Age":32,
	//	"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]
	// 	},
	// {
	//	"First":"Miss",
	//	"Last":"Moneypenny",
	//	"Age":27,
	//	"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]
	//	},
	//	{
	//	"First":"M",
	//	"Last":"Hmmmm",
	//	"Age":54,
	//	"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]
	// }]

	// instead of include os.Stdout, we could include a file
	// so that we could process the json inside, then we can output it to a file
	// or we could send it through an endpoint

	// we can use method channing
	// method channing is basically concateneting methods one after the other
	err = json.NewEncoder(os.Stdout).Encode(users[0].First)
	if err != nil {
		fmt.Println(err)
	}
	// "James"

}

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// next:
// https://www.youtube.com/watch?v=fZZclCKnr7k&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=123&ab_channel=AprendaGo
