// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"t1/utils"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := myMath.Average(xs)
	min, max := myMath.Extrema(xs)
	fmt.Printf("average of: %v is %.2f\n", xs, avg)
	fmt.Printf("min of: %v is %.2f\n", xs, min)
	fmt.Printf("max of: %v is %.2f\n", xs, max)
	fmt.Printf("That is all. Goodbye!\n")
}
