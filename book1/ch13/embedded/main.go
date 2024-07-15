package main

import (
	"fmt"
)

type Country struct {
	Name        string
	CapitalCity string
}

type Hotel struct {
	Name string
	Country
}

func main() {
	hotel := Hotel{
		Name: "Hotel super luxe",
		// without initialization CapitalCity is ""
		Country: Country{Name: "France"},
	}
	fmt.Println(hotel.Country.Name)
}
