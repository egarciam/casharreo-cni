package main

import (
	"fmt"
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

	input := make([]byte, 4096)

	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}

	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
