package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	go sub()

	go func() {
		fmt.Println("func() is running")
		time.Sleep(time.Second * 2)
		fmt.Println("func() is finished")
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("exit")
}

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub is finished")
}
