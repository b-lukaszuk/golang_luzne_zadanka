package main

import (
	"fmt"
)

func main() {
	b := "hello"
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])        // will output (ascii) code and not a letter
		fmt.Printf("%c\n", b[i]) // will output a letter
	}
}
