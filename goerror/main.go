package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 0)
	if err != nil {
		switch {
		case err == ErrDivideByZero:
			fmt.Println("Cannot divide by zero")
		default:
			fmt.Println("Unknown error")
		}
		return
	}

	fmt.Println("Result:", result)
}

var ErrDivideByZero = errors.New("Cannot divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero // return 0, errors.New("Cannot divide by zero")
		// return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
