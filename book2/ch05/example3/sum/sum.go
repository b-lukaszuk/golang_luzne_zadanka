package main

// []int is a slice of an array of any size
func Sum(numbers []int) int {
	sum := 0
	// range returns indx, elt of an array
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(vectsToSum ...[]int) []int {
	numberOfVectors := len(vectsToSum)
	// make - creates a slice with a starting capacity equal to numberOfVectors
	sums := make([]int, numberOfVectors)

	for i, numbers := range vectsToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
