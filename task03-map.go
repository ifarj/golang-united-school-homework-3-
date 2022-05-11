package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	arrKeys := make([]int, 0)
	for k := range input {
		arrKeys = append(arrKeys, k)
	}
	sort.Slice(arrKeys,  func(i, j int) bool { return input[i] < input[j] })
	for i := range arrKeys {
		result = append(result, input[arrKeys[i]])
	}
	return
}
