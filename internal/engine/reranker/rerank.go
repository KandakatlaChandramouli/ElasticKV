package reranker

type Candidate struct {
	ID    int
	Score float32
}

func Rerank(
	input []Candidate,
	boost float32,
) []Candidate {

	for i := range input {
		input[i].Score *= boost
	}

	return input
}
