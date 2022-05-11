package homework

import "sort"

func reverse(input []int64) []int64 {
	result := make([]int64, len(input))
	sort.Slice(input, func(i, j int) bool { return input[i] > input[j] })
	for i := range input {
		result[i] = input[i]
	}
	return result
}
