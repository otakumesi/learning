package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("q2.1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, "%d\n", 7)
	fmt.Fprintf(file, "%f\n", 0.05)
	fmt.Fprintf(file, "%s\n", "がんばっていこうや")
}
