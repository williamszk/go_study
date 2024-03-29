package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	strUrl := "ec2-3-84-186-251.compute-1.amazonaws.com:8000"
	// strUrl := "localhost:8000"
	conn, err := net.Dial("tcp", strUrl)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
