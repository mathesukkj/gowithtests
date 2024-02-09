package arraysslices

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			sums[i] += numbers[len(numbers)-1]
		}
	}
	return sums
}
