package main

import (
	"fmt"
	"os"
)

// based on exercise 1.2
func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
	}
}
