package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func getRandNum(minIncl uint, maxExcl uint) uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(int(maxExcl-minIncl)) + int(minIncl))
}

func getRandSpaceLine() string {
	spacelines := [...]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	var randInd uint = getRandNum(0, uint(len(spacelines)))
	return spacelines[randInd]
}

// in mln usd
func getRandPrice() uint {
	return getRandNum(36, 51)
}

// in km/sec
func getRandSpeed() uint {
	return getRandNum(16, 31)
}

func getTimeInDays(speedKmPerSec uint) uint {
	// in km from Earth
	const distToMars = 62100000
	var timeInDays float64 = float64(distToMars / speedKmPerSec / 3600 / 24)
	return uint(math.Ceil(timeInDays))
}

func getRandTripType() string {
	tripTypes := [2]string{"One-way", "Round-trip"}
	var randInd uint = getRandNum(0, uint(len(tripTypes)))
	return tripTypes[randInd]
}

func getTotalPrice(basePrice uint, days uint, tripType string) uint {
	totalPrice := float64(basePrice) * (1 + float64(days/20))
	if tripType == "Round-trip" {
		totalPrice *= 2
	}
	return uint(totalPrice)
}

func main() {
	fmt.Printf("Printing 10 random fake flights to Mars\n\n")
	fmt.Printf("%d\n", getTimeInDays(16))
	fmt.Printf("%d\n", getTimeInDays(30))
	fmt.Printf("\nThat's all. Goodbye!\n")
}
