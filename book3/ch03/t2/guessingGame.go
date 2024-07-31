package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printIntro() {
	fmt.Printf("===\n")
	fmt.Printf("Hello, let's play a game.\n")
	fmt.Printf("I'm thinking of a number between 1 and 100.\n")
	fmt.Printf("Can you guess it?\n")
	fmt.Printf("To make it easier after each guess I will give you a hint\n")
	fmt.Printf("by saying: 'too high' or 'too low'\n")
	fmt.Printf("Ready, let's start.\n")
	fmt.Printf("BTW. You should be able to abort ")
	fmt.Printf("the game by pressing Ctrl+C (perhaps twice) on your keyboard.\n")
	fmt.Printf("===\n")
}

func printOutro() {
	fmt.Printf("\n===\n")
	fmt.Printf("Game over. You won. Congratulations!")
	fmt.Printf("\n===\n")
}

func getGuess() uint {
	var guess uint
	fmt.Printf("\nEnter your guess (1-100) and press Enter: ")
	fmt.Scanln(&guess)
	return guess
}

func printHint(numToGuess, guess uint) {
	if numToGuess < guess {
		fmt.Printf("Too high.")
	} else if numToGuess > guess {
		fmt.Printf("Too low.")
	} else {
		fmt.Printf("You got it.\n")
	}

}

func runGame() {
	rand.Seed(time.Now().UnixNano())
	var secretNum uint = uint(rand.Intn(100) + 1)
	var guess uint = 0

	for {
		guess = getGuess()
		printHint(secretNum, guess)
		if guess == secretNum {
			break
		}
	}
}

func main() {
	printIntro()
	runGame()
	printOutro()
}
