package main

// arrays have a fix length
// if you pass an [4]int into a function that expects [5]int, it won't compile
func Sum(numbers [5]int) int {
	sum := 0
	// range returns indx, elt of an array
	for _, number := range numbers {
		sum += number
	}
	return sum
}
