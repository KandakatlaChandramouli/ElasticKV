package shard

import (
	"sync/atomic"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/storage"
)

type Manager struct {
	Workers    []*Worker
	ShardCount uint64
	Dropped    atomic.Uint64
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
) bool {

	shardID := Route(
		key,
		m.ShardCount,
	)

	req := WriteRequest{
		SequenceID: key,
		Payload:    payload,
	}

	ok := m.Workers[shardID].TryEnqueue(req)

	if !ok {
		m.Dropped.Add(1)
	}

	return ok
}

func (m *Manager) Stop() {

	for _, worker := range m.Workers {
		worker.Stop()
	}
}
