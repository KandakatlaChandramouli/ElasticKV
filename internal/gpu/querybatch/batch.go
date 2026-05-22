package querybatch

type Batch struct {
	Queries [][]float32
}

func Build(
	queries [][]float32,
	size int,
) []Batch {

	batches := make([]Batch, 0)

	for i := 0; i < len(queries); i += size {

		end := i + size

		if end > len(queries) {
			end = len(queries)
		}

		batches = append(
			batches,
			Batch{
				Queries: queries[i:end],
			},
		)
	}

	return batches
}
