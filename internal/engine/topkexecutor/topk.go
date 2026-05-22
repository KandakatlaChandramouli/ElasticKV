package topkexecutor

import "sort"

func TopK(
	values []float32,
	k int,
) []float32 {

	sort.Slice(
		values,
		func(i, j int) bool {
			return values[i] > values[j]
		},
	)

	if k > len(values) {
		k = len(values)
	}

	return values[:k]
}
