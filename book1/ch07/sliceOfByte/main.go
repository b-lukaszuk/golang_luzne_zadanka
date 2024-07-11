package main

import (
	"fmt"
)

func main() {
	b := make([]byte, 0)
	b = append(b, 255)
	// appending more than 255 will not compile (it produces an overflow error)
	// b = append(b, 256)
	b = append(b, 10)
	fmt.Println(b)
}
