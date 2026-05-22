package cardinality

func Estimate(
	rows int,
	selectivity float64,
) int {

	return int(
		float64(rows) * selectivity,
	)
}
