package main

import (
	"fmt"
	"time"
)

func main() {
	// we'll build an infinite loop which will wait each 1 second
	// this infinite loop will be looking for defined tasks
	fmt.Println("Sleeping 5 seconds...")
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleeping 5 seconds...")
	}
}
