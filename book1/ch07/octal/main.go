package main

import (
	"fmt"
)

func main() {
	n2 := 0x9f4 // hexadecimal
	fmt.Printf("Decimal: %d\n", n2)

	// n3 is represented the octal numeral system
	n3 := 02454 // 0 for octal
	// alternative: n3 := 0o2454
	fmt.Printf("decimal: %d\n", n3)

	n4 := 1324                                 // decimal
	fmt.Printf("octal: %o\n", n4)              // print in octal
	fmt.Printf("octal with prefix : %O\n", n4) // print in octal
}
