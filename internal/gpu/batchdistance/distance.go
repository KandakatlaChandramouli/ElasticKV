package batchdistance

func Compute(
	vectors [][]float32,
	query []float32,
) []float32 {

	scores := make([]float32, len(vectors))

	for i := range vectors {

		var sum float32

		for j := range query {
			sum += vectors[i][j] * query[j]
		}

		scores[i] = sum
	}

	return scores
}
