package main

import (
	"fmt"
	"struct/s1"
)

func main() {
	s1.Hello()
	type Person struct {
		firstName, lastName string
		age                 int
	}

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       25,
	}

	// create empty struct
	var p2 Person

	// create anonymous struct
	p3 := struct {
		name, title, company string
	}{
		name:    "John Doe",
		title:   "Software Engineer",
		company: "Google",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	// accessing struct fields
	fmt.Println("p1 firstname: ", p1.firstName)

	// create pointer to a struct variable
	ptr_p1 := &p1
	// fmt.Println("p1 firstname: ", (*ptr_p1).lastName) // access via pointer to struct
	fmt.Println("p1 firstname: ", ptr_p1.lastName) // access via pointer to struct, go automatically dereference the pointer

	// create struct using new keyword
	p4 := new(Person)
	p4.firstName = "Jane"
	p4.lastName = "Doe"
	p4.age = 30
	fmt.Println("Person 4 create with new keyword: ", *p4)
}
