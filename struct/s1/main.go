package s1

import "fmt"

func Comp() {
	type Human struct {
		name string
		age  int
	}

	type SuperHero struct {
		Human Human
		power string
	}

	clark := Human{"Clark Kent", 30}
	superman := SuperHero{Human: clark, power: "Super Strength"}

	fmt.Println("Name:", superman)
}

func main() {
	Comp()
	Hello()
}
