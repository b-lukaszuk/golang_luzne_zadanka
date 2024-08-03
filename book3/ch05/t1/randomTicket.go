package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
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
func getRandBasePrice() uint {
	return getRandNum(36, 51)
}

// in km/sec
func getRandSpeed() uint {
	return getRandNum(16, 31)
}

func getTimeInDays(speedKmPerSec uint) uint {
	const distToMars = 62.1e6 // km from Earth
	const secsPerDay = 60 * 60 * 24
	var timeInDays float64 = float64(distToMars / speedKmPerSec / secsPerDay)
	return uint(math.Ceil(timeInDays))
}

func getRandTripType() string {
	tripTypes := [2]string{"One-way", "Round-trip"}
	var randInd uint = getRandNum(0, uint(len(tripTypes)))
	return tripTypes[randInd]
}

func getTotalPrice(basePrice uint, days uint, tripType string) uint {
	// (1 + 1-days/100) because faster ships are more expensive
	totalPrice := float64(basePrice) * (1 + float64(1-days/100))
	if strings.Contains(tripType, "Round-trip") {
		totalPrice *= 2
	}
	return uint(totalPrice)
}

func getHeader() string {
	return `Spaceline           Days    Trip type       Price
=================================================
`
}

func getRandSingleTicket() string {
	spaceline := fmt.Sprintf("%-20s", getRandSpaceLine())
	daysNum := getTimeInDays(getRandSpeed())
	daysStr := fmt.Sprintf("%-8d", daysNum)
	tripType := fmt.Sprintf("%-16s", getRandTripType())
	basePrice := getRandBasePrice()
	totalPrice := fmt.Sprintf("$ %-3d",
		getTotalPrice(basePrice, daysNum, tripType))
	return fmt.Sprintf("%s%s%s%s\n", spaceline, daysStr, tripType, totalPrice)
}

func main() {
	fmt.Printf("Printing 10 random fake flights to Mars\n\n")
	fmt.Printf(getHeader())
	for i := 0; i < 10; i++ {
		fmt.Printf(getRandSingleTicket())
	}
	fmt.Printf("\nThat's all. Goodbye!\n")
}
