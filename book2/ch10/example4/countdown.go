package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprintln(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
