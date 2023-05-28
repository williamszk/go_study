package io

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// example of embedding interface
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriter2 interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type ReadWriter3 interface {
	Read(p []byte) (n int, err error)
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

func main() {

}
