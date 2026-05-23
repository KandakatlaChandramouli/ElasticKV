package batch

type VectorBatch struct {
	Vectors [][]float32
}

func New() *VectorBatch {

	return &VectorBatch{
		Vectors: [][]float32{},
	}
}

func (b *VectorBatch) Add(
	vec []float32,
) {

	b.Vectors = append(
		b.Vectors,
		vec,
	)
}
