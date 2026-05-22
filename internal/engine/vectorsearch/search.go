package vectorsearch

import "sort"

type Result struct {
	ID    int
	Score float32
}

func Search(
	scores []float32,
	topk int,
) []Result {

	results := make([]Result, len(scores))

	for i := range scores {
		results[i] = Result{
			ID:    i,
			Score: scores[i],
		}
	}

	sort.Slice(
		results,
		func(i, j int) bool {
			return results[i].Score > results[j].Score
		},
	)

	if topk > len(results) {
		topk = len(results)
	}

	return results[:topk]
}
