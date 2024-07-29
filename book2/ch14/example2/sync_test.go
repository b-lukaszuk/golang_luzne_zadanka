package main

import (
	"testing"
)

func assertCount(t testing.TB, got Counter, expectedCount int) {
	t.Helper()
	if got.Value() != expectedCount {
		t.Errorf("got %d, want %d", got.Value(), expectedCount)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3",
		func(t *testing.T) {
			counter := Counter{}
			counter.Inc()
			counter.Inc()
			counter.Inc()

			assertCount(t, counter, 3)
		})
}
