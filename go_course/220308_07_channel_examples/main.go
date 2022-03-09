package main

import "fmt"

func main() {
	chn01 := make(chan int)   // this is a bidirectional channel
	chn02 := make(<-chan int) // this is a sender channel
	chn03 := make(chan<- int) // this is a receiver channel

	fmt.Println("----------")
	fmt.Printf("chn01\t%T\n", chn01)
	fmt.Printf("chn02\t%T\n", chn02)
	fmt.Printf("chn03\t%T\n", chn03)

	// ----------
	// chn01   chan int
	// chn02   <-chan int
	// chn03   chan<- int

	// general to specific converts
	fmt.Println("-------------")
	fmt.Printf("chn01\t%T\n", (<-chan int)(chn01))
	fmt.Printf("chn01\t%T\n", (chan<- int)(chn01))

	// -------------
	// chn01   <-chan int
	// chn01   chan<- int

}

// next:
// https://www.youtube.com/watch?v=B1UArMoYDJ0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=148&ab_channel=AprendaGo
