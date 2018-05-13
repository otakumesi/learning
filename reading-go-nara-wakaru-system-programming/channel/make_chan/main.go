package main

import "fmt"

func main() {
	tasks := make(chan string, 10)
	go func() {
		tasks <- "Hello, "
		tasks <- "World"
		tasks <- "!!!"
		close(tasks)
	}()

	result := ""
	for task := range tasks {
		result += task
	}
	fmt.Println(result)
}
