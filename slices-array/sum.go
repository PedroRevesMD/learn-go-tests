package slicesarray

func Sum(values []int) int {
	sum := 0

	for i := range len(values) {
		sum += values[i]
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthSums := len(numbersToSum)
	sums := make([]int, lengthSums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
