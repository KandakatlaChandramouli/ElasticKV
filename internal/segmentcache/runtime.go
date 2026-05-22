package segmentcache

type Segment struct {
	ID uint64
}

type Runtime struct {
	Cache map[uint64]Segment
}

func NewRuntime() *Runtime {

	return &Runtime{
		Cache: make(map[uint64]Segment),
	}
}

func (r *Runtime) Put(
	id uint64,
) {

	r.Cache[id] = Segment{
		ID: id,
	}
}
