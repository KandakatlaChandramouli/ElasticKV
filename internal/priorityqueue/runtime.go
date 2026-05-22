package priorityqueue

type Entry struct {
	Priority uint64
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
	priority uint64,
) {

	r.Entries = append(
		r.Entries,
		Entry{
			Priority: priority,
		},
	)
}
