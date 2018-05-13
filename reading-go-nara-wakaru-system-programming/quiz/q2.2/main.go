package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{
		"hello",
		"world",
	})
	writer.Write([]string{
		"golang",
		"playground",
	})
	writer.Flush()
}
