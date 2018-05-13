package main

import (
	"io"
	"os"
)

func main() {
	// buffer := bufio.NewWriter(os.Stdout)
	// buffer.WriteString("test code")
	// buffer.Flush()
	// buffer.WriteString("example\n")
	// buffer.Flush()
	// fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
	// var buffer bytes.Buffer
	// encoder := json.NewEncoder(&buffer)
	// encoder.SetIndent("", "	")
	// encoder.Encode(map[string]string{
	// 	"test":  "code",
	// 	"hello": "world",
	// })
	// fmt.Printf("%v", buffer.String())

	file, err := os.Open("./notebook.md")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// size, err := io.ReadFull(file, buffer)
	// size, err := io.ReadFull(file, buffer)
	dst, err := os.Create("./copy.notebook.md")
	_, err = io.Copy(dst, file)
	if err != nil {
		panic(err)
	}
}
