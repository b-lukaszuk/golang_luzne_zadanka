package main

import (
	"fmt"
	"math/rand"
	"time"
)

func vacantRooms() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}

func main() {
	vacant := vacantRooms()
	fmt.Println("Vacant rooms:", vacant)
}
