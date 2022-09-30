package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/enceve/crypto/pad"
)

func main() {
	msg := "Hello1111111111111"
	argCount := len(os.Args[1:])
	blocksize := 16
	if argCount > 0 {
		msg = string(os.Args[1])
	}
	if argCount > 1 {
		blocksize, _ = strconv.Atoi(os.Args[2])
	}
	m := []byte(msg)
	fmt.Printf("Block size:\t%d bytes (%d bits)\n\n", blocksize, blocksize*8)
	pkcs7 := pad.NewPKCS7(blocksize)
	pad1 := pkcs7.Pad(m)
	fmt.Printf("PKCS7 Pad:\t%x", pad1)
	// this is a form of padding in which we repeat the number of the padding in the encoded form
	// for example if the block size is 8 bytes, and the last block uses 2, then we have 6 of padding
	// and it will be included as 060606060606
	res, _ := pkcs7.Unpad(pad1)
	fmt.Printf("\n Unpad:\t%s", res)
	x923 := pad.NewX923(blocksize)
	pad2 := x923.Pad(m)
	// in the case of X923 the padding is 00.00.00.00.00.06
	// for a case where we need to include 6 of padding, five 0s followed by a 6
	// 00.00.00.00.00.00.00.00.00.00.00.00.00.0e
	fmt.Printf("\n\nX923 Pad:\t%x", pad2)
	res, _ = x923.Unpad(pad2)
	fmt.Printf("\n Unpad:\t%s", res)
	ISO10126 := pad.NewISO10126(blocksize, nil)
	pad3 := ISO10126.Pad(m)
	fmt.Printf("\n\nXISO1012 Pad:\t%x", pad3)
	res, _ = ISO10126.Unpad(pad3)
	fmt.Printf("\n Unpad:\t%s", res)

	// the ISO 10126 includes a random stream of bytes but the last one
	// tells the number of random bytes for padding
}
