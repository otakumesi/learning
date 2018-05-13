package main

import (
	"flag"
	"io"
	"os"
)

var (
	dstOpt = flag.String("dst", "old.txt", "destination of copied file")
	srcOpt = flag.String("src", "new.txt", "source of copying file")
)

func main() {
	flag.Parse()

	file, err := os.Open(*srcOpt)
	if err != nil {
		panic(nil)
	}
	newFile, err := os.Create(*dstOpt)
	if err != nil {
		panic(err)
	}
	io.Copy(newFile, file)
}
