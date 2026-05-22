package simd

func Sum(
	data []uint64,
) uint64 {

	var total uint64

	for _, value := range data {

		total += value
	}

	return total
}
