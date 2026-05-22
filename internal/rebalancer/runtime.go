package rebalancer

type Shard struct {
	ID   uint64
	Node uint64
}

type Runtime struct {
	Shards []Shard
}

func NewRuntime() *Runtime {

	return &Runtime{
		Shards: make([]Shard, 0),
	}
}

func (r *Runtime) Move(
	id uint64,
	node uint64,
) {

	r.Shards = append(
		r.Shards,
		Shard{
			ID:   id,
			Node: node,
		},
	)
}
