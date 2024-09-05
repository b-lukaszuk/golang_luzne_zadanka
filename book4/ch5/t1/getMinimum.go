// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math"
)

func getMin(xs *[]int) int {
	var min int = math.MaxInt
	for _, value := range *xs {
		if value < min {
			min = value
		}
	}
	return min
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It finds the minimum of list of numbers.\n\n")

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Printf("The list: %v\n", x)
	fmt.Printf("The minimum: %d\n", getMin(&x))

	fmt.Println("\nThat's all. Goodbye!")
}
