package main

import (
	"fmt"
	"time"
)

func generator(ch chan int) {
	sum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		sum += i
	}
	ch <- sum
}

func main() {
	ch := make(chan int)

	go generator(ch)

	fmt.Println("Waiting for the result...")

	result := <-ch

	fmt.Println("Result:", result)
}
