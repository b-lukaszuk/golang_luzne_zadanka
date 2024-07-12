package main

import (
	"fmt"
)

func main() {
	const version string = "1.3.2" // typed constant
	fmt.Printf("version: %s\n", version)

	const version1 = "1.3.2" // untyped constant
	fmt.Printf("version: %s\n", version1)

	// untyped constant has no clear numeric type
	const occupancyLimit = 12

	var occupancyLimit1 uint8
	var occupancyLimit2 uint64
	var occupancyLimit3 float32

	occupancyLimit1 = occupancyLimit
	occupancyLimit2 = occupancyLimit
	occupancyLimit3 = occupancyLimit

	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)

	var occupancyLimit4 string
	// def type of occupancyLimit is int, it cannot be assigned to type string
	// occupancyLimit4 = occupancyLimit
	fmt.Println(occupancyLimit4)

	// constant defaults
	// default type is bool
	const isOpen = true
	// default type is rune (alias for int32)
	const myRune = 'r'
	// default type is int
	const occupLimit = 12
	// default type is float64
	const vatRate = 29.87
	// default type is complex128
	const complexNumber = 1 + 2i
	// default type is string
	const hotelName = "Gopher Hotel"
	fmt.Println(isOpen, myRune, occupLimit, vatRate, complexNumber, hotelName)
}
