// In this project we experiment with tranforming an integer into
// byte, little endian and big endian to see if there are any differences
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	tempBs4 := make([]byte, 4)

	myNum := 1

	fmt.Println("myNum:", myNum)
	fmt.Println("byte(myNum):", byte(myNum))

	myNum2 := 256

	fmt.Println("myNum2:", myNum2)
	fmt.Println("byte(myNum2):", byte(myNum2))

	binary.LittleEndian.PutUint32(tempBs4, uint32(myNum2))
	fmt.Println("tempBs4:", tempBs4)

	binary.BigEndian.PutUint32(tempBs4, uint32(myNum2))
	fmt.Println("big Endian tempBs4:", tempBs4)

	myNum3 := 1

	fmt.Println("myNum3:", myNum3)
	fmt.Println("byte(myNum3):", byte(myNum3))

	binary.LittleEndian.PutUint32(tempBs4, uint32(myNum3))
	fmt.Println("little Endian tempBs4:", tempBs4)

	binary.BigEndian.PutUint32(tempBs4, uint32(myNum3))
	fmt.Println("big Endian tempBs4:", tempBs4)

	// Conclusion: In the case of 1 byte, using just the type casting byte()
	// or using LittleEndian or BigEndian will not make a difference
	// but if we have more than 1 byte of space to store, we need to use
	// either LittleEndian or BigEndian

	// myArr := []int{1, 2, 3}
	// myArr := "hello there"
	// myArr = []byte(myArr)

}
