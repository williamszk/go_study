package main

import "io"

func main() {

}

// io.Writer is an interface, that means that we can pass any value with a type
// that implements the interface
// what are the types that implement the interface?
// ByteCounter implements the interface because ByteCounter has the method Write

// the return newWriter is also of a type io.Writer interface, that is, we can
// return any value of type that implements the io.Writer

// CountingWriter is a wrapper around io.Writer
func CountingWriter(w io.Writer) (newWriter io.Writer, numBytesWritten *int64) {
	w.Write()
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
