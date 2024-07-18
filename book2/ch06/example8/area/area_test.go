package area

import (
	"testing"
)

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, name string, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g, want %g", name, got, want)
		}
	}

	// slice of structs with two fields:
	// shape of type Shape
	// want of type float64
	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{10}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expectedArea: 36},
	}

	// to run a single test from the table type something like
	// go test -run TestArea/Rectangle
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.name, tt.shape, tt.expectedArea)
		})
	}
}
