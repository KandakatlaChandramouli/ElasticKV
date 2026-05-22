package truncation

type Runtime struct {
	Logs []uint64
}

func NewRuntime() *Runtime {

	return &Runtime{
		Logs: make([]uint64, 0),
	}
}

func (r *Runtime) Append(
	index uint64,
) {

	r.Logs = append(
		r.Logs,
		index,
	)
}

func (r *Runtime) Truncate(
	index uint64,
) {

	if int(index) >= len(r.Logs) {
		return
	}

	r.Logs = r.Logs[:index]
}
