// https://www.youtube.com/watch?v=cCWvFijhObU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=141&ab_channel=AprendaGo
package main

import (
	"fmt"
	"time"
)

// Make a race condition for a variable being incremented.
func main() {
	ctr := 0

	for i := 0; i < 100; i++ {
		go func() {
			hold := ctr
			time.Sleep(2)
			hold++
			ctr = hold
		}()
	}

	fmt.Println("Counter: ", ctr)
}

// how to build checking for race condition?
// go build -race
// go run main.go -race
