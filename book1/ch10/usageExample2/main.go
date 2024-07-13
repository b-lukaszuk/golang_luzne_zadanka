package main

import (
	"fmt"
)

func computePrice1(rate float32, nights int) (price float32) {
	price = rate * float32(nights) // price is declared in fn signature
	return                         // returns price (as declared in fn signature)
}

// price is declared in fn signature and by default initialized to 0
func computePrice2(rate float32, nights int) (price float32) {
	p := rate * float32(nights)
	p = p * 2
	return
}

func main() {
	johnPrice1 := computePrice1(145.90, 3)
	johnPrice2 := computePrice2(145.90, 3)
	fmt.Printf("TOTAL: %0.2f $\n", johnPrice1)
	fmt.Printf("TOTAL: %0.2f $\n", johnPrice2)
}
