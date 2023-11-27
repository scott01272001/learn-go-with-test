package arrayAndSlice

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numberLists ...[]int) []int {
	return []int{}
}
