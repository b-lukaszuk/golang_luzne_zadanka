// simple solution,
// not idiot/error-proof, not optimized for speed
package math

import (
	"math"
	"testing"
)

type testpair struct {
	values []float64
	min    float64
	max    float64
}

var tests = []testpair{
	{[]float64{0}, 0, 0},
	{[]float64{-2}, -2, -2},
	{[]float64{2}, 2, 2},
	{[]float64{1, 2}, 1, 2},
	{[]float64{1, 1, 1, 1, 1}, 1, 1},
	{[]float64{-1, -1, -1, -1, -1}, -1, -1},
	{[]float64{-3, 3, -1, -1, -2, 40, -2, -4, 0}, -4, 40},
}

func TestMinMax(t *testing.T) {
	for _, pair := range tests {
		low, high := GetMinAndMax(pair.values)
		if low != pair.min {
			t.Error("For", pair.values, "expected", pair.min, "got", low)
		}
		if high != pair.max {
			t.Error("For", pair.values, "expected", pair.max, "got", high)
		}
	}
}

func TestMinMaxEmptySlice(t *testing.T) {
	emptySlice := []float64{}
	low, high := GetMinAndMax(emptySlice)
	if !math.IsNaN(low) {
		t.Error("For", emptySlice, "expected", math.NaN(), "got", low)
	}
	if !math.IsNaN(high) {
		t.Error("For", emptySlice, "expected", math.NaN(), "got", high)
	}
}
