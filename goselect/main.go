package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		one <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		two <- "two"
	}()

	// select will block the code until one of the channels is ready
	// and execute that channel only
	select {
	case result := <-one:
		fmt.Println("Received:", result)

	case result := <-two:
		fmt.Println("Received:", result)

	default:
		fmt.Println("Default...")
		time.Sleep(200 * time.Millisecond)

	}

	fmt.Println("Done")

	close(one)
	close(two)
}
