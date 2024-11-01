// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"image"
	"time"
)

func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	delay := time.Millisecond * 1000
	next := time.After(delay)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current delay", delay)
			fmt.Println("current position is", pos)
			delay += time.Millisecond * 500
			next = time.After(delay)
		}
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It updates the time in worker by 0.5 sec by every move.\n")

	fmt.Println("\n==starting worker==")
	go worker()
	time.Sleep(time.Second * 5)
	fmt.Println("==worker stopped==")

	fmt.Printf("\nThat's all. Goodbye!\n")
}
