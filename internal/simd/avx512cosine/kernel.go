package avx512cosine

func Cosine(
	a []float32,
	b []float32,
) float32 {

	var dot float32
	var na float32
	var nb float32

	for i := range a {
		dot += a[i] * b[i]
		na += a[i] * a[i]
		nb += b[i] * b[i]
	}

	return dot / ((na * nb) + 1e-9)
}
