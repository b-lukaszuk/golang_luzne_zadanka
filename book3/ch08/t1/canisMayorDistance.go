package main

import (
	"fmt"
	"math/big"
)

func main() {
	canisMajorDistanceKm, _ := new(big.Float).SetString("236e15")
	lightSpeed := new(big.Float).SetFloat64(299792) // in km/sec
	secsInYear := new(big.Float).SetFloat64(60 * 60 * 24 * 365.2422)
	canisMajorDistanceLyr := new(big.Float)
	canisMajorDistanceLyr.Quo(canisMajorDistanceKm, lightSpeed)
	canisMajorDistanceLyr.Quo(canisMajorDistanceLyr, secsInYear)

	fmt.Printf("Toy program.\n")
	fmt.Printf("It estimates the distance from the Sun to Canis Major galaxy.\n\n")
	fmt.Printf("The galaxy Canis Major is ~%.2f [km] away from the Sun.\n", canisMajorDistanceKm)
	fmt.Printf("To travel this distance light needs ~%.2f [years].\n", canisMajorDistanceLyr)

	fmt.Printf("\nThat's all. Goodbye!\n")
}
