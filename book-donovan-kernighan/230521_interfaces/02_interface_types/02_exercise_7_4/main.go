package main

// Exercise 7.4: The strings.NewReader function returns a value that satisfies the io.Reader
// interface (and others) by reading from its argument, a string. Implement a simple version of
// NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string.

func main() {

	// strings.NewReader()
	// strings.Reader implements io.Reader

	// type Reader interface {
	// 	Read(p []byte) (n int, err error)
	// }

	// what the HTML parser in 5.2 does?

	// The NewReader creates a strings.Reader value
}
