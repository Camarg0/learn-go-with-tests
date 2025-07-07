package main

// Arrays are size encoded, if passed 4 elements it won't compile
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// slicesToSum is a slice of slices
func SumAll(slicesToSum ...[]int) []int {
	// // If I pass three slices, length's gonna be three
	// lengthOfSlices := len(slicesToSum)
	// // Create a new slice sized by the number os slices and with 0 values, for example lengthOfSlices = 2: sums = [0, 0]
	// sums := make([]int, lengthOfSlices)

	var sums []int

	for _, slice := range slicesToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
