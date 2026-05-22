package vectorindex

type Entry struct {
	Key uint64
}

type Runtime struct {
	Entries []Entry
}

func NewRuntime() *Runtime {

	return &Runtime{
		Entries: make([]Entry, 0),
	}
}

func (r *Runtime) Insert(
	key uint64,
) {

	r.Entries = append(
		r.Entries,
		Entry{
			Key: key,
		},
	)
}
