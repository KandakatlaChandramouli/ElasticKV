package productquantizer

type ProductQuantizer struct {
	Segments  int
	Centroids int
}

func NewPQ(
	segments int,
	centroids int,
) *ProductQuantizer {

	return &ProductQuantizer{
		Segments:  segments,
		Centroids: centroids,
	}
}

func (p *ProductQuantizer) Encode(
	vector []float32,
) []uint8 {

	codes := make([]uint8, len(vector))

	for i := range vector {
		codes[i] = uint8(int(vector[i]) % 255)
	}

	return codes
}
