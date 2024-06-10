package main

import "fmt"

func main() {
	// the defer keyword, which lets us postpones the execution of a function until the surrounding function returns.
	fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("main")
}
