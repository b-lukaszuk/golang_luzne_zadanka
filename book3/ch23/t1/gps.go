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

type gps struct {
	World           world
	CurrentLocation location
	Destination     location
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

func (g gps) getDistance() float64 {
	return g.World.getDistance(g.CurrentLocation, g.Destination)
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It displays the (hopefully correct) distance between two locations.\n\n")

	position := gps{
		World:           world{Name: "Mars", RadiusKm: 3389.5},
		CurrentLocation: location{Lat: -14.5684, Long: 175.472636},
		Destination:     location{Lat: -1.9462, Long: 354.4734},
	}

	fmt.Printf("The distance between %s and %s\n",
		position.CurrentLocation, position.Destination)
	fmt.Printf("on planet %s is equal to:\n", position.World.Name)
	fmt.Printf("%.2f km\n", position.getDistance())

	fmt.Printf("\nThat's all. Goodbye!\n")
}
