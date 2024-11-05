// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"time"
)

func mySleep(numOfSecs int64) {
	// time.Duration is in nanoseconds
	// 1 sec = 1000 ms, 1 ms = 1000 us, 1 us = 1000 ns
	<-time.After(time.Duration(numOfSecs * 1000 * 1000 * 1000))
}

func main() {

	fmt.Printf("Toy program.\n")
	fmt.Printf("It implements mySleep() using time.After().\n\n")

	for i := 1; i < 4; i++ {
		fmt.Println("Sleeping for", i, "seconds.")
		mySleep(int64(i))
	}

	fmt.Println("\nThat's all. Goodbye!")
}
