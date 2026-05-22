package pipeline

type Stage struct {
	ID uint64
}

type Runtime struct {
	Stages []Stage
}

func NewRuntime() *Runtime {

	return &Runtime{
		Stages: make([]Stage, 0),
	}
}

func (r *Runtime) Push(
	id uint64,
) {

	r.Stages = append(
		r.Stages,
		Stage{
			ID: id,
		},
	)
}
