package revisiontree

type Revision struct {
	ID uint64
}

type Runtime struct {
	Revisions []Revision
}

func NewRuntime() *Runtime {

	return &Runtime{
		Revisions: make([]Revision, 0),
	}
}

func (r *Runtime) Insert(
	id uint64,
) {

	r.Revisions = append(
		r.Revisions,
		Revision{
			ID: id,
		},
	)
}
