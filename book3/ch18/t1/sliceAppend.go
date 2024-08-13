// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It adds a few elts to slice.\n")
	fmt.Printf("Moreover, it prints its capacity whenever it changes.\n\n")

	sl := []int{0}
	prevCap := cap(sl)
	fmt.Printf("%v, capacity: %d\n", sl, prevCap)

	for i := 1; i < 20; i++ {
		fmt.Printf("adding num to slice\n")
		sl = append(sl, i)
		if prevCap != cap(sl) {
			// capacity seems to be doubled when running out of space
			prevCap = cap(sl)
			fmt.Printf("%v, length: %d, capacity: %d\n", sl, len(sl), prevCap)
		}
	}

	fmt.Printf("\nThat's all. Goodbye!\n")
}
