// run with: go test
package main

import (
	"testing"
)

// shorter version for arguments: got, want string
// achieves the same as got string, want string
func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper() // thx this, it will report line number of evoking fn when failed
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people",
		func(t *testing.T) {
			got := getGreeting("Chris", "")
			want := "Hello, Chris"
			assertCorrectMessage(t, got, want)

		})

	t.Run("saying 'Hello, World' when an empty string is supplied",
		func(t *testing.T) {
			got := getGreeting("", "")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
		})

	t.Run("in Spanish", func(t *testing.T) {
		got := getGreeting("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := getGreeting("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
