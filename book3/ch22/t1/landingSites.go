// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math"
)

type location struct {
	Lat  float64
	Long float64
}

func (l location) String() string {
	return fmt.Sprintf("[lat: %.2f, long: %.2f]", l.Lat, l.Long)
}

type world struct {
	Name     string
	RadiusKm float64
}

func getRadians(degs float64) float64 {
	return degs * math.Pi / 180.0
}

// distance calculation using the Spherical Law of Cosines
func (w world) getDistance(l1, l2 location) float64 {
	s1, c1 := math.Sincos(getRadians(l1.Lat))
	s2, c2 := math.Sincos(getRadians(l2.Lat))
	clong := math.Cos(getRadians(l1.Long - l2.Long))
	return w.RadiusKm * math.Acos(s1*s2+c1*c2*clong)
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It displays the (hopefully correct) distance between two locations.\n\n")

	planets := [...]world{
		// no need for world{Name: "Mars", RadiusKm: 3389.5},
		// due to the type declaration above
		{Name: "Mars", RadiusKm: 3389.5},
		{Name: "Venus", RadiusKm: 6051.8},
		{Name: "Earth", RadiusKm: 6371.0},
	}
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	fmt.Printf("The distance between %s and %s\n", spirit, opportunity)
	fmt.Printf("is equal to:\n")
	var dist float64
	for _, planet := range planets {
		dist = planet.getDistance(spirit, opportunity)
		fmt.Printf("\t- on %s = %.2f km\n", planet.Name, dist)
	}

	fmt.Printf("\nThat's all. Goodbye!\n")
}
