package main

import (
	"fmt"
)

func main() {
	fmt.Println("Beginning of program execution")
	// untyped constant has no limit
	// maximum value of an int is 9223372036854775807
	// 9223372036854775808 (max + 1 ) overflows int
	const profit = 9223372036854775808
	var profit2 int64 = profit // error, it overflows int
	fmt.Println(profit2)
	fmt.Println("End of program execution")
}
