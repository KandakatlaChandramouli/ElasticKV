package requestbatcher

type Request struct {
	Prompt string
}

type Batch struct {
	Requests []Request
}

func Build(
	requests []Request,
	batchSize int,
) []Batch {

	batches := make([]Batch, 0)

	for i := 0; i < len(requests); i += batchSize {

		end := i + batchSize

		if end > len(requests) {
			end = len(requests)
		}

		batches = append(
			batches,
			Batch{
				Requests: requests[i:end],
			},
		)
	}

	return batches
}
