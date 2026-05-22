package admission

type Runtime struct {
	Accepted uint64
	Rejected uint64
	Limit    uint64
}

func NewRuntime(
	limit uint64,
) *Runtime {

	return &Runtime{
		Limit: limit,
	}
}

func (r *Runtime) Allow(
	value uint64,
) bool {

	if value > r.Limit {
		r.Rejected++
		return false
	}

	r.Accepted++

	return true
}
