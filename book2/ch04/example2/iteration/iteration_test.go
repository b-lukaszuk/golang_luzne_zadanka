package iteration

import (
	"fmt"
	"testing"
)

func assertCorrectRepetition(t testing.TB, got string, want string) {
	t.Helper() // thx this, it will report line number of evoking fn when failed
	if got != want {
		t.Errorf("expected %q, but got %q", want, got)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("repeating string 0 times",
		func(t *testing.T) {
			repeated := Repeat("y", 0)
			want := ""
			assertCorrectRepetition(t, repeated, want)
		})

	t.Run("repeating string 3 times",
		func(t *testing.T) {
			repeated := Repeat("y", 3)
			want := "yyy"
			assertCorrectRepetition(t, repeated, want)
		})

	t.Run("repeating string 4 times",
		func(t *testing.T) {
			repeated := Repeat("yy", 4)
			want := "yyyyyyyy"
			assertCorrectRepetition(t, repeated, want)
		})
}

// run benchmarks with: go test -bench=.
// by default Benchmarks are run sequentially
func BenchmarkRepeat(b *testing.B) {
	// the framework determines the appropriate value of b.N
	// to have some decent results
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

// to see the output of ExampleRepeats
// run tests in verbose mode: go test -v
func ExampleRepeat() {
	repeated := Repeat("x", 2)
	fmt.Println(repeated)
	// Output: xx
}

func ExampleRepeat2() {
	repeated := Repeat("xx", 2)
	fmt.Println(repeated)
	// Output: xxxx
}
