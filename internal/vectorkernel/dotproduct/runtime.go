package dotproduct

func Compute(
	a []float32,
	b []float32,
) float32 {

	var result float32

	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}

	return result
}
