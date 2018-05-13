package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	a := io.ReaderAt(programming, 5)
	sc := io.LimitReader(io.MultiReader(system, computer), 1)
	i := io.NewSectionReader(programming, 8, 1)
	stream := io.MultiReader(a, sc, computer)

	// ASCII
	io.Copy(os.Stdout, stream)
}
