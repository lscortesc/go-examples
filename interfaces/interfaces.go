package main

import (
	"fmt"
	"os"
)

// Reader Interfaces are implicitly implemented
type Reader interface {
	Read(b []byte) (n int, err error)
}

// Writer interface
type Writer interface {
	Write(b []byte) (n int, err error)
}

// ReaderWriter interface
type ReaderWriter interface {
	Reader
	Writer
}

func interfaces() {
	var w Writer

	// os.Stdout implements writer
	w = os.Stdout

	fmt.Fprintf(w, "Hola, writer\n")
}
