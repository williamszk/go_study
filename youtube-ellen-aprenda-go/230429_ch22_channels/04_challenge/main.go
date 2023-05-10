// https://www.youtube.com/watch?v=eoA4jVD7RQE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=160&ab_channel=AprendaGo

package main

import "fmt"

func main() {

	var c = make(chan int)

	for i := 0; i < 10; i++ {
		go feedTheChannel(c)
	}

	for j := 0; j < 100; j++ {
		fmt.Println(j, <-c)
	}

}

func feedTheChannel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
