package transport

import (
	"sync"
	"sync/atomic"
	"time"
)

type BatchReplicator struct {
	Client      *Client
	Queue       chan ReplicationFrame
	BatchSize   int
	FlushPeriod time.Duration
	Running     atomic.Bool
	WG          sync.WaitGroup
	Sent        atomic.Uint64
	Dropped     atomic.Uint64
	Batches     atomic.Uint64
}

func NewBatchReplicator(
	client *Client,
	queueSize int,
	batchSize int,
	flushPeriod time.Duration,
) *BatchReplicator {

	r := &BatchReplicator{
		Client:      client,
		Queue:       make(chan ReplicationFrame, queueSize),
		BatchSize:   batchSize,
		FlushPeriod: flushPeriod,
	}

	r.Running.Store(true)

	return r
}

func (r *BatchReplicator) Start() {

	r.WG.Add(1)

	go func() {

		defer r.WG.Done()

		ticker := time.NewTicker(
			r.FlushPeriod,
		)

		defer ticker.Stop()

		batch := make([]ReplicationFrame, 0, r.BatchSize)

		flush := func() {

			for _, frame := range batch {

				err := r.Client.Send(
					frame,
				)

				if err != nil {
					continue
				}

				r.Sent.Add(1)
			}

			if len(batch) > 0 {
				r.Batches.Add(1)
			}

			batch = batch[:0]
		}

		for {

			select {

			case frame, ok := <-r.Queue:

				if !ok {

					flush()

					return
				}

				batch = append(
					batch,
					frame,
				)

				if len(batch) >= r.BatchSize {
					flush()
				}

			case <-ticker.C:

				flush()
			}
		}
	}()
}

func (r *BatchReplicator) Replicate(
	frame ReplicationFrame,
) bool {

	select {

	case r.Queue <- frame:
		return true

	default:
		r.Dropped.Add(1)
		return false
	}
}

func (r *BatchReplicator) Stop() {

	if !r.Running.Load() {
		return
	}

	r.Running.Store(false)

	close(r.Queue)

	r.WG.Wait()
}
