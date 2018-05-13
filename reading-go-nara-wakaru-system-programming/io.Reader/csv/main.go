package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `13101,"100 ","1000003","トウキョウト","東京都"
13101,"101 ","1010003","トウキョウト","東京都"
13101,"100 ","1000012","トウキョウト","東京都"`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[2:4])
	}
}
