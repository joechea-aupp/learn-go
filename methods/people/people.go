package people

import "fmt"

type Person struct {
	FirstName, LastName string
	Dob                 string
}

func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) UpdateFirstName(newName string) {
	p.FirstName = newName
}
