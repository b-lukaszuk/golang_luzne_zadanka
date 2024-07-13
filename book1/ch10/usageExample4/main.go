package main

import (
	"fmt"
)

// Go got hoisting, but I prefer to declare function before usage
func printHeader() {
	fmt.Println("Hotel Golang")
	fmt.Println("San Francisco, CA")
}

func main() {
	printHeader()
}
