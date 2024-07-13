package main

import (
	"fmt"
)

// errors inside
func printHeader() {
	// HotelName is undefined, it is in main, but printHeader does not see it
	fmt.Println("Hotel ", HotelName)
	fmt.Println("San Francisco, CA")
	// cannot add return, because there is no return in the signature
}

func main() {
	const HotelName = "Golang"
	printHeader()
}
