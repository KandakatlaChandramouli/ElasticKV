package snapshotdelta

type Delta struct {
	Key uint64
}

type Runtime struct {
	Deltas []Delta
}

func NewRuntime() *Runtime {

	return &Runtime{
		Deltas: make([]Delta, 0),
	}
}

func (r *Runtime) Append(
	key uint64,
) {

	r.Deltas = append(
		r.Deltas,
		Delta{
			Key: key,
		},
	)
}
