// type: go test -cover
// to get the covarage of your tests
package main

import (
	"reflect"
	"testing"
)

func assertCorrectSum(t testing.TB, got int, want int, nums []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, nums)
	}
}

// assigning a function to a variable
// here the shortcut := wouldn't work,
// but it should inside TestSumAllTails
var assertCorrectSumAll = func(t testing.TB, got []int, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assertCorrectSumAll(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertCorrectSumAll(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertCorrectSumAll(t, got, want)
	})
}
