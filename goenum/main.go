package main

import (
	"fmt"
)

type DayOfTheWeek uint8

const (
	Monday DayOfTheWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Printf("Monday: %d\n", Monday)
	fmt.Printf("Wednesday: %d\n", Wednesday)
	fmt.Printf("Friday: %d\n", Friday)
}
