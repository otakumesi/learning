package main

import (
	"fmt"
	"strings"
)

var source = "123 1.2341.0e4 test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var a, b float64
	var str string
	fmt.Fscan(reader, &i, &a, &b, &str)
	fmt.Printf("%d, %f, %f, %s", i, a, b, str)
}
