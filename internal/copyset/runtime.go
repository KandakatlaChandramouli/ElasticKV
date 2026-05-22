package copyset

type Replica struct {
	Node uint64
}

type Runtime struct {
	Replicas []Replica
}

func NewRuntime() *Runtime {

	return &Runtime{
		Replicas: make([]Replica, 0),
	}
}

func (r *Runtime) Add(
	node uint64,
) {

	r.Replicas = append(
		r.Replicas,
		Replica{
			Node: node,
		},
	)
}
