package main

import "fmt"

func main() {
	var t interface{}
	t = 10

	switch t.(type) {
	case int:
		fmt.Println("t is an int")
	case string:
		fmt.Println("t is a string")
	case float64:
		fmt.Println("t is a float64")
	case bool:
		fmt.Println("t is a bool")
	default:
		fmt.Println("t is another type")

	}
}
