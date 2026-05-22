package mutationlog

type Entry struct {
	ID uint64
}

type Runtime struct {
	Log []Entry
}

func NewRuntime() *Runtime {

	return &Runtime{
		Log: make([]Entry, 0),
	}
}

func (r *Runtime) Append(
	id uint64,
) {

	r.Log = append(
		r.Log,
		Entry{
			ID: id,
		},
	)
}
