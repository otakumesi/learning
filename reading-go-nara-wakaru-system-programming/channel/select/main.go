package main

import (
	"fmt"
	"time"
)

func main() {
	task := doTask()
	exit := returnExit()
	for {
		select {
		case data := <-task:
			fmt.Println(data)
		case <-exit:
			break
		}
	}
}

func returnExit() chan int {
	exit := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		exit <- 1
	}()
	return exit
}

func doTask() chan int {
	task := make(chan int)

	go func() {
		task <- 1
	}()
	return task
}
