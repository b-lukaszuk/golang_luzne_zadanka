// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

func getFib(n uint) uint {
	if n < 2 {
		return n
	}
	return getFib(n-1) + getFib(n-2)
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It generates numbers from the Fibbonaci sequence.\n\n")

	for i := uint(0); i < 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, getFib(i))
	}

	fmt.Println("\nThat's all. Goodbye!")
}
