package vectorpredicate

func Similarity(
	score float32,
	threshold float32,
) bool {

	return score >= threshold
}
