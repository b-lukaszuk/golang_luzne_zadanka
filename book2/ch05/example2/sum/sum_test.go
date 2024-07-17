// type: go test -cover
// to get the covarage of your tests
package main

import (
	"testing"
)

func assertCorrectSum(t testing.TB, got int, want int, nums []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, nums)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectSum(t, got, want, numbers)

	})

	t.Run("collection of size 3", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectSum(t, got, want, numbers)
	})
}
