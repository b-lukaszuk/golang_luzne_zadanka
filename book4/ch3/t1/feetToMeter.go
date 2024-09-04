// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandNumber(minIncl float64, maxExcl float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(maxExcl-minIncl) + minIncl
}

func getMeters(feet float64) float64 {
	return feet * 0.3048
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It converts (hopefully correctly) feet to meters.\n\n")

	var feet float64
	for i := 0; i < 10; i++ {
		feet = getRandNumber(1, 100)
		fmt.Printf("%.4f feet is %.4f meter.\n", feet, getMeters(feet))
	}

	fmt.Println("\nThat's all. Goodbye!")
}
