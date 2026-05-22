package hybridsearch

type Result struct {
	ID           int
	VectorScore  float32
	KeywordScore float32
}

func Combine(
	vector float32,
	keyword float32,
) float32 {

	return (vector * 0.7) +
		(keyword * 0.3)
}
