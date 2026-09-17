package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestTee(t *testing.T) {
	reader := strings.NewReader("123")
	readerOut, readerIn, _ := os.Pipe()
	go func() { io.Copy(readerIn, reader); readerIn.Close() }()
	os.Stdin = readerOut
	os.Args = []string{"123", "1.txt"}
	main()
}
