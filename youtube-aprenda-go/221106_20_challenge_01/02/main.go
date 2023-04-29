package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	createGoRoutines(10)
	wg.Wait()
}

func createGoRoutines(qtd int) {
	wg.Add(qtd)
	for j := 0; j < qtd; j++ {
		i := j
		go func(i int) {
			fmt.Println("I ran goroutine number:", i)
			wg.Done()
		}(i)
	}
}
