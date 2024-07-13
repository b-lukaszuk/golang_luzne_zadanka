package main

import (
	"fmt"
)

func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}

func main() {
	johnPrice := computePrice(145.90, 3)
	paulPrice := computePrice(145.90, 3)
	robPrice := computePrice(145.90, 3)

	total := johnPrice + paulPrice + robPrice
	fmt.Printf("TOTAL: %0.2f $\n", total)
}
