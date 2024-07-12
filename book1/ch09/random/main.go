package main

import (
	"fmt"
	"math/rand"
)

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel", hotelName)
	var roomsAvailable int
	// without seeding rand.Intn(100) will produce the same number every time
	var rooms, roomsOccupied int = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupied
	fmt.Println(roomsAvailable, "rooms available")
}
