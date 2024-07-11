package main

import (
	"fmt"
)

func main() {
	n := 2548                       // decimal
	n2 := 0x9f4                     // hexadecimal
	fmt.Printf("%x\n", n)           // prints in hex format, lower case lettters
	fmt.Printf("%X\n", n)           // prints in hex format, upper case lettters
	fmt.Printf("%x\n", n2)          // prints in hex format, lower case lettters
	fmt.Printf("Decimal: %d\n", n2) // prints in decimal format
}
