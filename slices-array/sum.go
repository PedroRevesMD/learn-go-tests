package slicesarray

func Sum(values []int) int {
	sum := 0

	for i := range len(values) {
		sum += values[i]
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
