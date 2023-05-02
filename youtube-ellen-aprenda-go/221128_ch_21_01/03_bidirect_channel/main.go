package main

import "fmt"

func main() {
	myChan := make(chan int)
	myChanRcv := make(<-chan int)
	myChanSnd := make(chan<- int)

	fmt.Printf("--- Channel and its types ----------\n")
	fmt.Printf("myChan   \t%T\n", myChan)
	fmt.Printf("myChanRcv\t%T\n", myChanRcv)
	fmt.Printf("myChanSnd\t%T\n", myChanSnd)

	fmt.Printf("--- Convert bidirect channel to unidirect channel ----------\n")
	fmt.Printf("myChan   \t%T\n", (<-chan int)(myChan))
	fmt.Printf("myChan   \t%T\n", (chan<- int)(myChan))

}
