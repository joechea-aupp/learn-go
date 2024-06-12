package main

import "fmt"

func speak(arg string, ch chan string) {
	// add variable of arg to channel ch
	ch <- arg
}

func main() {
	// initialize channel ch with type of string
	// channel is identify as chan keyword
	// by default channel can take only one cargo, if you want to take more than one cargo, you can specify the number of cargo in the channel by adding number after chan keyword
	ch := make(chan string, 2) // this channel can take 2 cargo

	// we call go routine by specify go keyword infront of a function.
	// go routine is a lightweight thread of execution
	// this will store string into a channel ch
	go speak("Hello", ch)
	go speak("Hi", ch)

	// now it is time to retrieve data from channel ch and put it into variable data
	data1 := <-ch
	fmt.Println("data 1:", data1)

	data2, ok := <-ch
	fmt.Println("data 2:", data2, ok) // ok is false if the channel is closed

	close(ch)
}
