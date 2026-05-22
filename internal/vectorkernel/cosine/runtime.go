package cosine

func Similarity(
	a []float32,
	b []float32,
) float32 {

	var dot float32
	var normA float32
	var normB float32

	for i := 0; i < len(a); i++ {

		dot += a[i] * b[i]

		normA += a[i] * a[i]

		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dot
}
