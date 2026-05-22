package consensusqueue

type Entry struct {
	Term uint64
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
	term uint64,
) {

	r.Entries = append(
		r.Entries,
		Entry{
			Term: term,
		},
	)
}
