package main

// What happens when we create a buffered channel and no one read from it?
// But someone will input values into it.
// The program will continue as normal.
func main() {

	myChan := make(chan int, 2)

	myChan <- 10

}
