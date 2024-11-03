// simple solution,
// not idiot/error-proof, not optimized for speed
package math

import "math"

func GetMinAndMax(xs []float64) (float64, float64) {
	if len(xs) == 0 {
		return math.NaN(), math.NaN()
	}
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, x := range xs {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}
