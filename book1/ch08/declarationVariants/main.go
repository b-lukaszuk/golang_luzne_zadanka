package main

import (
	"fmt"
)

func main() {
	var rn, fn int
	fmt.Println(rn, fn)

	var pswd string
	fmt.Printf("password '%s'\n", pswd)

	var roomNumber, floorNumber int = 154, 3
	fmt.Println(roomNumber, floorNumber)

	var password = "notSecured"
	fmt.Println(password)

	// WARNING: short variable declaration cannot be used outside functions
	roomNumber2, floorNumber2 := 151, 2 // short variable declaration
	fmt.Println(roomNumber2, floorNumber2)

	roomNumber1 := 154 // short variable declaration
	fmt.Println(roomNumber1)
}
