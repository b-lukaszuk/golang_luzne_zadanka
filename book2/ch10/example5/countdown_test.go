package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountdown(t *testing.T) {
	fmt.Println("Running TestCountdown, this may take a while")
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	fmt.Println("Finished running TestCountdown")
}
