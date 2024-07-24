package main

import (
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "https://www.facebook.com"
	fastURL := "https://www.quii.dev"

	// the comparison of external urls response times in unreliable
	// the test may sometimes pass, and sometimes not
	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
