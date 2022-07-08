package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	//slice2 := make([]int64, len(input))
	//copy := copy(slice2, input)
	reversed := make([]int64, 0)

	for i := len(input) - 1; i >= 0; i-- {
		reversed = append(reversed, input[i])
	}
	return reversed
}
