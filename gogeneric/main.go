package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//	func sum[T any](a, b T) T {
//		fmt.Println(a, b)
//		return b
//	}
func sum[T constraints.Ordered](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum("1", "2"))
}
