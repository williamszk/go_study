// https://www.youtube.com/watch?v=4jXSU2jw3Ag&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=126&ab_channel=AprendaGo

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())

	waitGroup.Add(2)
	go func1() // this is a go routine
	go func2()

	waitGroup.Wait()
	fmt.Println(runtime.NumCPU())
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1", i)
		time.Sleep(1)
	}

	waitGroup.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2", i)
		time.Sleep(50)
	}

	waitGroup.Done()
}
