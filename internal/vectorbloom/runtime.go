package vectorbloom

func Contains(
	data []uint64,
	target uint64,
) bool {

	for _, value := range data {

		if value == target {
			return true
		}
	}

	return false
}
