package main

import "fmt"

func main() {
	a := 10
	var p *int = &a
	// p := &a
	fmt.Println("address: ", p)
	fmt.Println("value: ", *p)

	*p = 20
	fn1(&p)
	fmt.Println("value: ", a)
}

func fn1(arg **int) {
	**arg++
}
