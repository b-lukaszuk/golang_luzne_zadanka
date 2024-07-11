package main

import (
	"fmt"
)

func main() {
	var aRune rune = 'Z'
	fmt.Printf("Unicode Code point of '%c' is %U\n", aRune, aRune)
}
