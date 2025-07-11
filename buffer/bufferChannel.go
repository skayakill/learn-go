package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	messages := make(chan int, 3)
	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	} ()
	for i := 0; i < 5; i++ {
		fmt.Println("send data", 1)
		messages <- 1
	}
}