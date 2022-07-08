package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	sortedByKeys := []string{}
	sortedKeys := []int{}
	for k := range input {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)
	for i := 0; i < len(sortedKeys); i++ {
		sortedByKeys = append(sortedByKeys, input[sortedKeys[i]])
	}
	return sortedByKeys
}
