package allocator

type Runtime struct {
	Buffers [][]byte
}

func NewRuntime() *Runtime {

	return &Runtime{
		Buffers: make([][]byte, 0),
	}
}

func (r *Runtime) Allocate(
	size int,
) []byte {

	buffer := make(
		[]byte,
		size,
	)

	r.Buffers = append(
		r.Buffers,
		buffer,
	)

	return buffer
}
