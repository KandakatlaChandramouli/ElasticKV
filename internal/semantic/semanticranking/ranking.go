package semanticranking

import "sort"

type Result struct {
	ID    int
	Score float32
}

func Rank(
	results []Result,
) []Result {

	sort.Slice(
		results,
		func(i, j int) bool {
			return results[i].Score > results[j].Score
		},
	)

	return results
}
