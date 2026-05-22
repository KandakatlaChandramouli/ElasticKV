package remotevector

type Vector struct {
	ID     int
	Values []float32
}

func Replicate(
	vectors []Vector,
) int {

	return len(vectors)
}
