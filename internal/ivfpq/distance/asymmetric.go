package distance

func AsymmetricDistance(
	query []float32,
	encoded []uint8,
) float32 {

	var score float32

	for i := range query {
		score += query[i] * float32(encoded[i])
	}

	return score
}
