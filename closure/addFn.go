package main

import "fmt"

func main() {
	add := myFn()
	add(5)

	fmt.Println(
		add(10), add(3))
}

func myFn() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v
		return sum
	}
}
