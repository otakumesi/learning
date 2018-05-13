package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5)
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("size='%d' input='%s'\n", n, buffer[:n])
	}
}
