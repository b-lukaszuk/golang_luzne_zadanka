// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getPerimeter() float64
}

type Circle struct {
	centerX, centerY, radius float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func getDistance(x1, y1, x2, y2 float64) float64 {
	dxSquared := math.Pow(x1-x2, 2)
	dySquared := math.Pow(y1-y2, 2)
	return math.Sqrt(dxSquared + dySquared)
}

func (r *Rectangle) getLength() float64 {
	return getDistance(r.x1, r.y1, r.x1, r.y2)
}

func (r *Rectangle) getWidth() float64 {
	return getDistance(r.x1, r.y1, r.x2, r.y1)
}

func (r *Rectangle) getPerimeter() float64 {
	return 2*r.getLength() + 2*r.getWidth()
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle (x1=%.2f, y1=%.2f, x2=%.2f, y2=%.2f)",
		r.x1, r.y1, r.x2, r.y2)
}

func (c *Circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle (cx=%.2f, cy=%.2f, r=%.2f)",
		c.centerX, c.centerY, c.radius)
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It calculates (hopefully correctly) perimeters of shapes.\n\n")

	shapes := []Shape{&Circle{0, 0, 10}, &Rectangle{0, 0, 10, 10}}
	for _, s := range shapes {
		fmt.Printf("%v, perimeter: %.2f\n", s, s.getPerimeter())
	}

	fmt.Println("\nThat's all. Goodbye!")
}
