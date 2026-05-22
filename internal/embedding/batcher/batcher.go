package batcher

type Batch struct {
	Inputs []string
}

func Build(
	inputs []string,
	batchSize int,
) []Batch {

	batches := make([]Batch, 0)

	for i := 0; i < len(inputs); i += batchSize {

		end := i + batchSize

		if end > len(inputs) {
			end = len(inputs)
		}

		batches = append(
			batches,
			Batch{
				Inputs: inputs[i:end],
			},
		)
	}

	return batches
}
