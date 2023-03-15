package math

func Average(input []float64) float64 {
	sum := 0.0
	for _,v := range input {
		sum += v;
	}
	length := float64(len(input))
	return sum/length
}