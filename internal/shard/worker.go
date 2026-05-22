package shard

import (
	"sync"
	"sync/atomic"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/storage"
)

type WriteRequest struct {
	SequenceID uint64
	Payload    []byte
}

type Worker struct {
	ID        uint64
	Engine    *storage.Engine
	Queue     chan WriteRequest
	WG        sync.WaitGroup
	Running   bool
	Dropped   atomic.Uint64
	Processed atomic.Uint64
}

func NewWorker(
	id uint64,
	engine *storage.Engine,
	queueSize int,
) *Worker {

	return &Worker{
		ID:      id,
		Engine:  engine,
		Queue:   make(chan WriteRequest, queueSize),
		Running: true,
	}
}

func (w *Worker) Start() {

	w.WG.Add(1)

	go func() {

		defer w.WG.Done()

		for req := range w.Queue {

			_, err := w.Engine.Write(
				req.SequenceID,
				req.Payload,
			)

			if err != nil {
				continue
			}

			w.Processed.Add(1)
		}
	}()
}

func (w *Worker) TryEnqueue(
	req WriteRequest,
) bool {

	select {

	case w.Queue <- req:
		return true

	default:
		w.Dropped.Add(1)
		return false
	}
}

func (w *Worker) Stop() {

	if !w.Running {
		return
	}

	w.Running = false

	close(w.Queue)

	w.WG.Wait()
}
