package vectorscan

func Scan(
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
