package recovery

import (
	"github.com/KandakatlaChandramouli/ElasticKV/internal/wal"
)

type Runtime struct {
	Manager *wal.Manager
}

func NewRuntime(
	manager *wal.Manager,
) *Runtime {

	return &Runtime{
		Manager: manager,
	}
}

func (r *Runtime) Replay() (
	map[uint64][]byte,
	error,
) {

	state := make(
		map[uint64][]byte,
	)

	err := r.Manager.Replay(
		func(entry wal.Entry) error {

			state[entry.SequenceID] =
				entry.Payload

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return state, nil
}
