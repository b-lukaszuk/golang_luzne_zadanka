package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var ageJohn int = rand.Intn(10)
	fmt.Println("John is", ageJohn, "years old")

	for i := 0; i < ageJohn; i++ {
		fmt.Println("iteration N", i)
	}
}
