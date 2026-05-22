package snapshotcache

type Snapshot struct {
	ID uint64
}

type Runtime struct {
	Cache map[uint64]Snapshot
}

func NewRuntime() *Runtime {

	return &Runtime{
		Cache: make(map[uint64]Snapshot),
	}
}

func (r *Runtime) Put(
	id uint64,
) {

	r.Cache[id] = Snapshot{
		ID: id,
	}
}
