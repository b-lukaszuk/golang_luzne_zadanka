// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	var tmpX int = *x
	*x = *y
	*y = tmpX
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It swaps two integers.\n\n")

	x := 1
	y := 2
	fmt.Printf("x = %d, y = %d.\n", x, y)
	fmt.Printf("Swapping x with y.\n")
	swap(&x, &y)
	fmt.Printf("x = %d, y = %d.\n", x, y)
	fmt.Printf("Swapping x with y again.\n")
	swap(&y, &x)
	fmt.Printf("x = %d, y = %d.\n", x, y)

	fmt.Println("\nThat's all. Goodbye!")
}
