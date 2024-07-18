package area

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// method
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
