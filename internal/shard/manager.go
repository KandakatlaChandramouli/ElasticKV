package shard

import (
	"github.com/KandakatlaChandramouli/ElasticKV/internal/storage"
)

type Manager struct {
	Workers    []*Worker
	ShardCount uint64
}

func NewManager(
	shardCount uint64,
	engineSize int,
	queueSize int,
) (*Manager, error) {

	workers := make([]*Worker, shardCount)

	for i := uint64(0); i < shardCount; i++ {

		engine, err := storage.Open(
			"shard_"+string(rune(i+'0'))+".db",
			engineSize,
		)

		if err != nil {
			return nil, err
		}

		worker := NewWorker(
			i,
			engine,
			queueSize,
		)

		worker.Start()

		workers[i] = worker
	}

	return &Manager{
		Workers:    workers,
		ShardCount: shardCount,
	}, nil
}

func (m *Manager) Dispatch(
	key uint64,
	payload []byte,
) {

	shardID := Route(
		key,
		m.ShardCount,
	)

	req := WriteRequest{
		SequenceID: key,
		Payload:    payload,
	}

	m.Workers[shardID].Queue <- req
}

func (m *Manager) Stop() {

	for _, worker := range m.Workers {
		worker.Stop()
	}
}
