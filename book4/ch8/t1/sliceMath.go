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
	fmt.Printf("average of: %v is %.2f", xs, avg)
}
