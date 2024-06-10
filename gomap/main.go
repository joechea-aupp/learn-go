package main

import "fmt"

func main() {
	// declare map with string as key and int as value
	m := map[string]int{
		"a": 0,
		"b": 1,
	}

	fmt.Println("first map:", m)

	// map with struct
	type User struct {
		Name string
		id   int
	}
	mperson := map[string]User{
		"john": {"John", 1},
		"doe":  {"Doe", 2},
	}

	fmt.Println("second map with struct:", mperson)

	// map with for loop
	for key, value := range mperson {
		result := fmt.Sprintf("key: %s, value: %v\n", key, value)
		fmt.Println(result)
	}

	// retrieve value from map
	john := mperson["john"]
	fmt.Println("john:", john)

	// when retrieve key that is not exist, will return empty struct with zero value
	empty := mperson["empty"]
	fmt.Println("empty:", empty)

	// check if key exist
	_, ok := mperson["john"]
	fmt.Println("john exist:", ok)

	_, ok = mperson["empty"]
	fmt.Println("empty exist:", ok)

	// delete key from map
	delete(mperson, "john")
	fmt.Println("after delete john:", mperson)

	// update value
	mperson["doe"] = User{"Doe", 3}
	fmt.Println("after update doe:", mperson)
}
