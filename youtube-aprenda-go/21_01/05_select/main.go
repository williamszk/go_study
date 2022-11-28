// https://www.youtube.com/watch?v=dp8s5jAc7h0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=149&ab_channel=AprendaGo

package main

func main() {

	theChan01 := make(chan int)
	theChan02 := make(chan int)

	totalBoth := 500
	halfTotal := totalBoth / 2

	go func(total int) {
		for i := 0; i < total; i++ {

		}

	}(halfTotal)

}
