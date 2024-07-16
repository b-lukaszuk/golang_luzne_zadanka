// run with: go test
package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := getHelloWorld()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
