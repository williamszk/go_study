// https://www.youtube.com/watch?v=Y9QEvz4D_9E&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=139&ab_channel=AprendaGo
package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {

	waitGroup.Add(2)

	go myGoroutine01()
	go myGoroutine02()

	waitGroup.Wait()
}

func myGoroutine01() {
	fmt.Println("We ran the myGoroutine01")

	waitGroup.Done()
}

func myGoroutine02() {
	fmt.Println("We ran the myGoroutine02")
	waitGroup.Done()
}
