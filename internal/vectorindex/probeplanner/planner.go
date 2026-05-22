package probeplanner

func Plan(
	clusters int,
	probes int,
) []int {

	result := make([]int, 0)

	for i := 0; i < probes && i < clusters; i++ {
		result = append(result, i)
	}

	return result
}
