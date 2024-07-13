package main

import (
	"fmt"
)

func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}

func main() {
	fmt.Printf("Rare %0.2f, nights %d, price %0.2f\n", 44.0, 4, computePrice(44, 4))
}
