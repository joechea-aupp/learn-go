package main

import "fmt"

func main() {
	fruits := [3]string{"apple", "banana", "cherry"}

	// getting array information with len and cap
	fmt.Println("length:", len(fruits), "capacity:", cap(fruits))

	for i, fruit := range fruits {
		// i, fruit are values that destructured from range fruits
		// which is index and the actual value
		result := fmt.Sprintf("Index: %d, Fruit: %s", i, fruit)

		fmt.Println(result)
	}

	// multi-demensional array
	// matrix := [2][2]int{{1, 2}, {3, 4}}
	// the ... is used to let the compiler determine the size of the array
	matrix := [...][2]int{{1, 2}, {3, 4}}

	for _, row := range matrix {
		fmt.Println("Row", row)
	}

	// slice
	programmingLanguages := []string{"Go", "Java", "Python"}
	fmt.Println("length:", len(programmingLanguages), "capacity:", cap(programmingLanguages))
	for i, language := range programmingLanguages {
		fmt.Println(i, language)
	}

	// multi-demensional slice
	matrixSlice := [][]int{{1, 2}, {3, 4}}
	for _, row := range matrixSlice {
		fmt.Println("Row", row)
	}
}
