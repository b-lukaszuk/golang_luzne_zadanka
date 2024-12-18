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
	var sums []int
	for _, numbers := range vectsToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(vectsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range vectsToSum {
		// slices can be sliced, slice[low:high]
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))

		}
	}
	return sums
}
