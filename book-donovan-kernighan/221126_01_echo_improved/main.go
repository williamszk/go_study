package main

import (
	"fmt"
	"os"
	"strings"
)

// this is the third implementation of echo
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
