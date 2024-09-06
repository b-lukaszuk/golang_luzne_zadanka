// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

func makeOddGenerator() func() uint {
	var i uint = 1
	return func() uint {
		var result uint = i
		i += 2
		return result
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It generates consecutive odd numbers.\n\n")

	getNextOdd := makeOddGenerator()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", getNextOdd())
	}

	fmt.Println("\nThat's all. Goodbye!")
}
