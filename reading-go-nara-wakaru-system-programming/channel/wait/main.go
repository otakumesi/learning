package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("staring sub()")
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Finished!!!")
}
