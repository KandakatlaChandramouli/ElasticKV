package batchscheduler

type Batch struct {
	ID uint64
}

type Runtime struct {
	Batches []Batch
}

func NewRuntime() *Runtime {

	return &Runtime{
		Batches: make([]Batch, 0),
	}
}

func (r *Runtime) Schedule(
	id uint64,
) {

	r.Batches = append(
		r.Batches,
		Batch{
			ID: id,
		},
	)
}
