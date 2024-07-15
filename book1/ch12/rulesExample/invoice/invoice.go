package invoice

import (
	"fmt"
	"rulesExample/currency"
)

func init() {
	fmt.Println("Running invoice init")
}

func Print() {
	fmt.Println("INVOICE Number X")
	fmt.Println(54, currency.EuroSymbol())
}
