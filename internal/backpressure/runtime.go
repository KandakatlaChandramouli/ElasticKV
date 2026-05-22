package backpressure

type Runtime struct {
	Depth uint64
	Limit uint64
}

func NewRuntime(
	limit uint64,
) *Runtime {

	return &Runtime{
		Limit: limit,
	}
}

func (r *Runtime) Push() bool {

	r.Depth++

	return r.Depth <= r.Limit
}
