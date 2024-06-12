package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int32
	B string
}

func reflectType() {
	var a int32 = 42
	var b string = "forty two"

	typeA := reflect.TypeOf(a)
	fmt.Println(typeA)

	typeB := reflect.TypeOf(b)
	fmt.Println(typeB)

	t := T{A: 1, B: "hello"}
	typeT := reflect.TypeOf(t)
	fmt.Println(typeT)

	// numfield is use to get number of field that is available in struct
	fmt.Println("numfield:", typeT.NumField())

	for i := 0; i < typeT.NumField(); i++ {
		// field is use to get field information from struct
		// it will return struct field information like name, type, tag etc
		field := typeT.Field(i)
		fmt.Println(field)
		fmt.Println(field.Name, field.Type)
	}
}

type Adder interface {
	Add(int, int) int
}

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func typeWithStruct() {
	var ptrAdder *Adder
	adderType := reflect.TypeOf(ptrAdder).Elem()

	c := Calculator{}
	calcType := reflect.TypeOf(c)
	calcTypePtr := reflect.TypeOf(&c)

	fmt.Println(calcType, "calcpointer:", calcTypePtr, "adderType:", adderType, calcTypePtr.Implements(adderType))
}

func main() {
	typeWithStruct()
	// reflectType()
}
