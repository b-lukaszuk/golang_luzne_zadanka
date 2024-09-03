// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It answers a few questions.\n\n")

	fmt.Printf("Q1. What's the largest 8-digit binary?\n")
	fmt.Printf("A1. It is: (2^8 - 1) D = %.0f D.\n", math.Pow(2, 8))
	binary8digit := "1111 1111"
	decimal, _ := strconv.ParseInt(strings.ReplaceAll(binary8digit, " ", ""), 2, 64)
	fmt.Printf("Let's test it: %s B = %d D.\n\n", binary8digit, decimal)

	fmt.Printf("Q2. What's 32,132 x 42,452?\n")
	fmt.Printf("A2. 32,132 x 42,452 = %d.\n\n", 32_132*42_452)

	exampleStr := "Example string"
	fmt.Printf("Q3. What's the length of '%s'?\n", exampleStr)
	fmt.Printf("A3. The length of that string is %d.\n", len(exampleStr))

	fmt.Println("\nThat's all. Goodbye!")
}
