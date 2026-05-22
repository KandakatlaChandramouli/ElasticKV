package topk

import "sort"

func Select(
	values []float32,
	k int,
) []float32 {

	sorted := make(
		[]float32,
		len(values),
	)

	copy(
		sorted,
		values,
	)

	sort.Slice(
		sorted,
		func(i, j int) bool {
			return sorted[i] > sorted[j]
		},
	)

	return sorted[:k]
}
