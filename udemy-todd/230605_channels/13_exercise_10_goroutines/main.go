package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func(iArg int) {
				for j := 0; j < 10; j++ {
					c <- iArg*100 + j
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
