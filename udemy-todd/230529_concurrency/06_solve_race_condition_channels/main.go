// I failed to solve the race condition using only channels
package main

import (
	"fmt"
	"runtime"
)

// create a race condition

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var myChan = make(chan int)
	// counter := 0
	// myChan <- counter

	// let's use a semaphore
	var done = make(chan bool)

	const gs = 100
	// var wg sync.WaitGroup
	// wg.Add(gs)
	// I think that with channels I don't need to use wait groups
	// the channel will stop execution of code when we need to send or
	// receive data from it

	for i := 0; i < gs; i++ {
		go func() {
			// counter := counter
			counter := <-myChan
			// we need a way for the goruntime to do somehting else
			// time.Sleep(time.Second*2) // this is one way
			runtime.Gosched() // this is another way

			counter++
			// counter = v
			myChan <- counter

			// wg.Done()
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
			done <- true // this signals that the go routine is done
		}()
	}

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// this for loop will make sure we have finished all goroutines
	// that are publishing into the channel, before we can close the channel
	for i := 0; i < gs; i++ {
		<-done
	}

	// print the final result
	fmt.Println("Final result:", <-myChan)

	close(myChan)

	// wg.Wait()
	// fmt.Println("counter\t", counter)
	// for v := range myChan {
	// 	fmt.Println("Final result:", v)
	// }
}
