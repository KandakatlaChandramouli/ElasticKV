package ringbuffer

type Runtime struct {
	Buffer []uint64
	Head   uint64
	Size   uint64
}

func NewRuntime(
	size uint64,
) *Runtime {

	return &Runtime{
		Buffer: make(
			[]uint64,
			size,
		),
		Size: size,
	}
}

func (r *Runtime) Push(
	value uint64,
) {

	index := r.Head % r.Size

	r.Buffer[index] = value

	r.Head++
}
