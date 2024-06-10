package main

import "fmt"

type MyInterface interface {
	Method()
}

type MyInt struct {
	property int
}

func (m MyInt) Method() {}

func main() {
	var i MyInterface
	i = MyInt{property: 1}

	fmt.Printf("(%v, %T)\n", i, i) // Output: ({10}, main.MyType)
}
