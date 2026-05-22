package l2

func Distance(
	a []float32,
	b []float32,
) float32 {

	var distance float32

	for i := 0; i < len(a); i++ {

		diff := a[i] - b[i]

		distance += diff * diff
	}

	return distance
}
