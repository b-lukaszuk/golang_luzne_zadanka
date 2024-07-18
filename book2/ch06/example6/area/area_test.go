package area

import (
	"testing"
)

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	// below: table driven tests

	// slice of structs with two fields:
	// shape of type Shape
	// want of type float64
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// structs are anonymus, because fields are not named explicitly
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
