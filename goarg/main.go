package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	argsWithProg := os.Args
	numA, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		log.Println("error:", err)
		os.Exit(2)
	}

	numB, err := strconv.Atoi(argsWithProg[2])
	if err != nil {
		log.Println("error:", err)
		os.Exit(2)
	}

	result := numA + numB
	log.Println("Result:", result)
}

func example1() {
	// os.Args store argments from os in a slice
	argsWithProg := os.Args
	log.Println(argsWithProg)
	// access the arguments from index 1
	argsWithoutProg := os.Args[1:]
	log.Println(argsWithoutProg)
	// access the argument at index 3
	arg := os.Args[3]
	log.Println(arg)
}
