package main

import (
	"fmt"
)

// main func lives in its own goroutine
func main() {
	fmt.Println("01. Main Goroutine starts here")
	// channel is a medium of communication and synchronization between defferent goroutines
	c := make(chan string, 4) // buffer of 4

	// this anonymous func runs in a different goroutine than the main func
	// it creates and sends a string to the main goroutine
	fmt.Println("02. Anonymous func Goroutine starts here")
	go func(input chan string) {
		input <- "hello from the Anonymous func Goroutine 1"
		input <- "hello from the Anonymous func Goroutine 2"
		input <- "hello from the Anonymous func Goroutine 3"
		input <- "hello from the Anonymous func Goroutine 4"
	}(c)
	fmt.Println("03. Anonymous func Goroutine sends message here")

	// greeting var receives the message from the anonymous func goroutine
	// it will block the flow of the main goroutine until it receives a message
	greeting := <-c
	fmt.Println("04. Message received")
	for greeting = range c {
		fmt.Println(greeting)
	}
	fmt.Println("05. Message prints")
	// and prints it to the main go routine
}

func helloWorld() {
	fmt.Println("Hello world!")
}
