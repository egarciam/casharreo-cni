package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin.
type FooReader struct{}

// Read data from stdin
func (FooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in>")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (FooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out>")
	return os.Stdout.Write(b)
}

func main() {
	// instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}
