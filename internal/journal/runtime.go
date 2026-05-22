package journal

type Runtime struct {
	Entries [][]byte
}

func NewRuntime() *Runtime {

	return &Runtime{
		Entries: make([][]byte, 0),
	}
}

func (r *Runtime) Append(
	data []byte,
) {

	r.Entries = append(
		r.Entries,
		data,
	)
}
