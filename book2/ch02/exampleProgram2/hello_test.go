// run with: go test
package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := getGreeting("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
