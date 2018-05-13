package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// header := bytes.NewBufferString("----- HEADER -----\n")
	// content := bytes.NewBufferString("Example of io.MultiReader\n")
	// footer := bytes.NewBufferString("----- FOOTER -----\n")

	header := strings.NewReader("----- HEADER -----\n")
	content := strings.NewReader("Example of io.MultiReader\n")
	footer := strings.NewReader("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}
