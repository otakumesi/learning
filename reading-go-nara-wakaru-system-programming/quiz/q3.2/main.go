package main

import (
	"crypto/rand"
	"os"
)

func main() {
	file, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1024)
	rand.Reader.Read(buffer)
	file.Write(buffer)
}
