package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	writer, err := os.Create("./hello.zip")
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	strWriter, err := zipWriter.Create("./hello.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(strWriter, strings.NewReader("Hello, World"))
}
