package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <filePath>\n", os.Args[0])
		return
	}
	out, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("Open file error: %v\n", err)
	}
	defer out.Close()
	_, err = io.Copy(io.MultiWriter(out, os.Stdout), os.Stdin)
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
		return
	}
}
