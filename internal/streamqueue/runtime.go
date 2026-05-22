package streamqueue

type Entry struct {
	ID uint64
}

type Runtime struct {
	Entries []Entry
}

func NewRuntime() *Runtime {

	return &Runtime{
		Entries: make([]Entry, 0),
	}
}

func (r *Runtime) Push(
	id uint64,
) {

	r.Entries = append(
		r.Entries,
		Entry{
			ID: id,
		},
	)
}
