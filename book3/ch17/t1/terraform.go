// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

type Planet string
type Planets []Planet

func (ps Planets) terraform() {
	for i, p := range ps {
		ps[i] = Planet(fmt.Sprintf("New %s", p))
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It 'terraforms' planets from a slice.\n\n")

	planets := Planets{"Mercury",
		"Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune"}
	// it seems that in Go one cannot just take an array for indexing
	// (as they would, e.g. in Python), like
	// planets2 := planets[[3, 6, 7]]
	// or
	// planets2 := planets[[]int{3, 6, 7}]
	// which is inconvenient and kinda lame
	planets2 := planets[6:]
	planets3 := planets[3:4]
	// one could create an empty slice of Planets
	// fill it (in for loop) with append() with planets[i]
	// then terraform it,
	// but the planets could be unaffected (a copy made on the way?)

	fmt.Printf("The planets are %s\n", planets)
	planets2.terraform()
	planets3.terraform()
	fmt.Printf("Terraforming some planets, result: %s\n", planets)
	// to better display planets, could use strings package like
	// strings.Join(data, ", "), but first data needs to be converted
	// to an array of strings

	fmt.Printf("\nThat's all. Goodbye!\n")
}
