package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	waitGroup.Add(2)

	go func1()
	go func2()

	waitGroup.Wait()

}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1", i)
		time.Sleep(1)
	}

	waitGroup.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2", i)
		time.Sleep(1)
	}

	waitGroup.Done()
}
