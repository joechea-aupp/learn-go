package main

import (
	"fmt"
	"gomethod/people"
)

func main() {
	employee := people.Person{
		FirstName: "John",
		LastName:  "Doe",
		Dob:       "01/01/2000",
	}

	employee.UpdateFirstName("Jane")
	fmt.Println(employee.GetFullName())
}
