// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It prints integers (1-100) evenly divisible by 3.\n\n")

	for i := 3; i < 101; i += 3 {
		fmt.Printf("%d\n", i)
	}

	fmt.Println("\nThat's all. Goodbye!")
}
