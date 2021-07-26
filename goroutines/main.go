package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {
	go display("hello")
	// go display("more hello's")

	display("world yoooo")

	fmt.Println("Main function")

	go func() {
		fmt.Println("Anonymous goroutine")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Main function")
}
