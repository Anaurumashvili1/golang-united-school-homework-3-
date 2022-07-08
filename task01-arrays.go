package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	sum := float32(0)

	for i := 0; i < 15; i++ {
		sum += input[i]
	}
	return sum / 15
}
