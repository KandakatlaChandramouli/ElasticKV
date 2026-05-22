package vectorclock

type Clock struct {
	Node uint64
	Time uint64
}

type Runtime struct {
	Clocks map[uint64]uint64
}

func NewRuntime() *Runtime {

	return &Runtime{
		Clocks: make(map[uint64]uint64),
	}
}

func (r *Runtime) Tick(
	node uint64,
) {

	r.Clocks[node]++
}
