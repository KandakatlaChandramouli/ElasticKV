package vectorizedscan

func Scan(
	column []int,
	predicate func(int) bool,
) []int {

	result := make([]int, 0)

	for _, value := range column {

		if predicate(value) {
			result = append(
				result,
				value,
			)
		}
	}

	return result
}
