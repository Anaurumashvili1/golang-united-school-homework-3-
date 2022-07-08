package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	reversed := make([]int64, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		for j := 0; j < len(input); j++ {
			reversed[j] = input[i]
		}
	}
	return reversed
}
