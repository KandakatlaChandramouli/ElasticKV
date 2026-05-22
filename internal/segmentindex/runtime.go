package segmentindex

type Segment struct {
	ID uint64
}

type Runtime struct {
	Segments map[uint64]Segment
}

func NewRuntime() *Runtime {

	return &Runtime{
		Segments: make(map[uint64]Segment),
	}
}

func (r *Runtime) Put(
	id uint64,
) {

	r.Segments[id] = Segment{
		ID: id,
	}
}
