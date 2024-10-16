package perimeter

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}

// will not compile, Go does not have function overloading
func Area(circ Circle) float64 {
	return 0
}
