package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandNum(minIncl uint, maxExcl uint) uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(int(maxExcl-minIncl)) + int(minIncl))
}

// in cents
func getRandCoin() uint {
	coins := [...]uint{5, 10, 25}
	var randInd uint = getRandNum(0, uint(len(coins)))
	return coins[randInd]
}

func main() {
	var piggyBank uint = 0       // in cents
	var curBalance float64 = 0.0 // in dollars
	var curCoin uint = 0         // in cents
	// lowered it to $2, printout was too long
	var upperLimit uint = 20 * 10 // in cents

	fmt.Printf("Toy program.\nIt adds coins at random to a piggy bank.\n")
	fmt.Printf("Initial balance is $%.2f.\n", float64(piggyBank))
	fmt.Printf("Final balance is >= $%.2f.\n\n", float64(upperLimit)/100)

	for {
		curCoin = getRandCoin()
		piggyBank += curCoin
		curBalance = float64(piggyBank) / 100
		fmt.Printf("Adding %d cents to the balance\n", curCoin)
		fmt.Printf("Piggy bank. Current balance $%.2f\n", curBalance)
		if piggyBank > upperLimit {
			break
		}
	}

	fmt.Printf("\nPiggy bank. Target balance reached/exceeded.\n\n")
	fmt.Printf("\nThat's all. Goodbye!\n")
}
