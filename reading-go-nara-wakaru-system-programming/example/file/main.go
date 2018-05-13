package main

import (
	"io"
	"os"
)

func main() {
	open()
	read()
}

func open() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Hello, File System")
}

func read() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
