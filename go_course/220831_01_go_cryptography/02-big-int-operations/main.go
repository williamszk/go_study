package main

// https://medium.com/asecuritysite-when-bob-met-alice/golang-and-cryptography-914db9d7069f

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func isHexString(s string) bool {
	return strings.ContainsAny(s, "abc")
}

func toBigInt(s string) *big.Int {

	// x, _ := new(big.Int).SetString("0", 16)
	var x *big.Int

	if strings.HasPrefix(s, "0x") {
		s = strings.Replace(s, "0x", "", 1)
		x, _ = new(big.Int).SetString(s, 16)
	} else if strings.HasPrefix(s, "0") {
		x, _ = new(big.Int).SetString(s, 8)
	} else if isHexString(s) {
		x, _ = new(big.Int).SetString(s, 16)
	} else {
		x, _ = new(big.Int).SetString(s, 10)
	}

	return (x)
}

// What are v1, v2 and p1?
//  - v1 // an integer value 1
//  - v2 // an integer value 2
//  - p1 // a prime number used in the mathemtical module operation
func main() {
	v1 := "2e63d239b3fc47433f19789843e30514b53dfd8773ebc915c0bc774e5368dbb6"
	v2 := "2e63d239b3fc47433f19789843e30514b53dfd8773ebc915c0bc774e5368dbb4"
	p1 := "127"
	argCount := len(os.Args[1:])
	//	x2, _ := new(big.Int).SetString(v1, 16)
	if argCount > 0 {
		v1 = os.Args[1]
	}
	if argCount > 1 {
		v2 = os.Args[2]
	}
	if argCount > 2 {
		p1 = os.Args[3]
	}
	x1 := toBigInt(v1)
	x2 := toBigInt(v2)
	prime := toBigInt(p1)
	fmt.Printf("Value 1 in hex: %x\n", v1)
	fmt.Printf("Value 1 in decimal: %s\n", v1)
	fmt.Printf("Value 2 in hex: %x\n", v1)
	fmt.Printf("Value 2 in decimal: %s\n", v1)
	fmt.Printf("v1+v2= %s\n", new(big.Int).Add(x1, x2).String())
	fmt.Printf("v1-v2= %s\n", new(big.Int).Sub(x1, x2).String())
	fmt.Printf("v1*v2= %s\n", new(big.Int).Mul(x1, x2).String())

	fmt.Printf("p = %s\n", prime)
	result := new(big.Int).Exp(x1, x2, prime) // x1^x2 (mod prime)
	fmt.Printf("v1^v2 (mod p)= %s\n", result)

	result = new(big.Int).Mul(x1, x2)
	result = new(big.Int).Mod(result, prime)
	fmt.Printf("v1*v2 (mod p)= %s\n", result)

	result = new(big.Int).Add(x1, x2)
	result = new(big.Int).Mod(result, prime)
	fmt.Printf("v1+v2 (mod p)= %s\n", result)

	Invx2 := new(big.Int).ModInverse(x2, prime) // x2^{-1} (mod prime)
	result = new(big.Int).Mul(x1, Invx2)
	result = new(big.Int).Mod(result, prime)
	fmt.Printf("v1/v2 (mod p)= %s\n", result)
}
