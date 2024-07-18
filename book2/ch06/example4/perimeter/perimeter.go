package perimeter

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// method
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
