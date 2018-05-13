package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	copyN(os.Stdout, strings.NewReader("Hello, World!!!"), 5)
}

func copyN(dst io.Writer, src io.Reader, length int) {
	reader := io.LimitReader(src, int64(length))
	io.Copy(dst, reader)
}
