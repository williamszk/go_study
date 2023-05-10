// https://www.youtube.com/watch?v=o3C_Cy2bn7Q&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=159&ab_channel=AprendaGo
package main

import "fmt"

func main() {

	myChan := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			myChan <- i
		}
		close(myChan)
	}()

	for val := range myChan {
		fmt.Println(val)
	}

}
