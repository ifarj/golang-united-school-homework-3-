package homework

func reverse(input []int64) []int64 {
	length := len(input)
	result := make([]int64, length)
	for i := range input {
		result[i] = input[length-1-i]
	}
	return result
}
