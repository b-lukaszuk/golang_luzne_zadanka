package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")

	// switch expression
	switch ageJohn {
	case 10:
		fmt.Println("john is 10 years old")
	case 20:
		fmt.Println("john is 20 years old")
	case 100:
		fmt.Println("john is 100 years old")
	default:
		fmt.Println("John is neither 10, 20 nor 100 years old")
	}

	// switch statement; expression
	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40:
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	default:
		fmt.Println("ageSum of Jonh and Paul is something else than 10, 20, 30 or 40")
	}

	// switch (without statement, without expression)
	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	default:
		fmt.Println("agePaul == ageJohn")
	}
}
