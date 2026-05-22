package iterator

type Runtime struct {
	Data []uint64
}

func NewRuntime(
	data []uint64,
) *Runtime {

	return &Runtime{
		Data: data,
	}
}

func (r *Runtime) Iterate(
	fn func(uint64),
) {

	for _, value := range r.Data {

		fn(value)
	}
}
