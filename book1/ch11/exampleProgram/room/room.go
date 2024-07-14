package room

import (
	"fmt"
)

const roomText = "room %d: %d people / %d nights\n"

// to export a function from a package it needs to start with a Capital Letter
func PrintDetails(roomNumber int, size int, nights int) {
	fmt.Printf(roomText, roomNumber, size, nights)
}
