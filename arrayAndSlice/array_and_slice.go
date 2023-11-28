package arrayAndSlice

import "fmt"

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numberLists ...[]int) []int {
	sums := make([]int, len(numberLists))
	for i, numbers := range numberLists {
		sums[i] = Sum(numbers)
	}
	fmt.Println(sums)
	return sums
}

func SumAllTails(numberLists ...[]int) []int {
	sums := make([]int, len(numberLists))
	for i, numbers := range numberLists {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(numbers[1:])
		}
	}
	return sums
}
