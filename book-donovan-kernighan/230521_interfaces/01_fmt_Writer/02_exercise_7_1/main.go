package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	aText := "Using the ideas from ByteCounter, implement counters for words and for lines. You will find bufio.ScanWords useful."
	var c WordCounter = 0
	c.Write([]byte(aText))
	fmt.Println(c)

	aText = "Using the ideas"
	var c2 WordCounter = 0
	c2.Write([]byte(aText))
	fmt.Println(c2)
}

type WordCounter int

// use:
// bufio.ScanWords

// type Writer interface {
//     Write(p []byte) (n int, err error)
// }

// Write will return the number of words.
func (c *WordCounter) Write(p []byte) (n int, err error) {

	reader := bytes.NewReader(p)
	// `reader` is of type `*bytes.Reader`, and it seems that it implements
	// the interface `io.Reader`
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		scanner.Text()
		*c += 1
	}

	return int(*c), nil
}

// create a function that is like fmt.Println or fmt.Sprintf which will use the
// function fmt.Fprintf and it will return the number of words inside the string
// passed.
