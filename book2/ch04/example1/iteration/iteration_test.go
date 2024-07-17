package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// run benchmarks with: go test -bench=.
// by default Benchmarks are run sequentially
func BenchmarkRepeat(b *testing.B) {
	// the framework determines the appropriate value of b.N
	// to have some decent results
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
