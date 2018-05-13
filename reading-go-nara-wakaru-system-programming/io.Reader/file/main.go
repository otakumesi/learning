package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, file)
}
