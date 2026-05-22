package segment

type Segment struct {
	ID      uint64
	Entries uint64
}

type Runtime struct {
	Segments []Segment
}

func NewRuntime() *Runtime {

	return &Runtime{
		Segments: make([]Segment, 0),
	}
}

func (r *Runtime) Create(
	id uint64,
	entries uint64,
) {

	r.Segments = append(
		r.Segments,
		Segment{
			ID:      id,
			Entries: entries,
		},
	)
}
