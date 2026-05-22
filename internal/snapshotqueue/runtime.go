package snapshotqueue

type Snapshot struct {
	ID uint64
}

type Runtime struct {
	Queue []Snapshot
}

func NewRuntime() *Runtime {

	return &Runtime{
		Queue: make([]Snapshot, 0),
	}
}

func (r *Runtime) Push(
	id uint64,
) {

	r.Queue = append(
		r.Queue,
		Snapshot{
			ID: id,
		},
	)
}
