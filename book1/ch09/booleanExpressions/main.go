package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)

	fmt.Println("rooms:", rooms, "\b, roomsOccupied:", roomsOccupied)

	// Example 1: is there more rooms than roomsOccupied
	fmt.Println("boolean expression: 'rooms > roomsOccupied' is equal to:")
	fmt.Println(rooms > roomsOccupied)

	// Example 2: is there more roomsOccupied than rooms
	fmt.Println("boolean expression: 'roomsOccupied > rooms' is equal to:")
	fmt.Println(roomsOccupied > rooms)

	// Example 3: is roomsOccupied equal to rooms
	fmt.Println("boolean expression: 'roomsOccupied > rooms' is equal to:")
	fmt.Println(roomsOccupied == rooms)
}
