package transport

import (
	"bufio"
	"sync"
	"sync/atomic"
	"time"
)

type WireBatchReplicator struct {
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

func NewWireBatchReplicator(
	client *Client,
	queueSize int,
	batchSize int,
	flushPeriod time.Duration,
) *WireBatchReplicator {

	r := &WireBatchReplicator{
		Client:      client,
		Queue:       make(chan ReplicationFrame, queueSize),
		BatchSize:   batchSize,
		FlushPeriod: flushPeriod,
	}

	r.Running.Store(true)

	return r
}

func (r *WireBatchReplicator) Start() {

	r.WG.Add(1)

	go func() {

		defer r.WG.Done()

		ticker := time.NewTicker(
			r.FlushPeriod,
		)

		defer ticker.Stop()

		batch := make(
			[]ReplicationFrame,
			0,
			r.BatchSize,
		)

		flush := func() {

			if len(batch) == 0 {
				return
			}

			buffer := AcquireBuffer()

			writer := bufio.NewWriterSize(
				buffer,
				2*1024*1024,
			)

			success := uint64(0)

			for _, frame := range batch {

				err := WriteFrame(
					writer,
					frame,
				)

				if err != nil {
					continue
				}

				success++
			}

			err := writer.Flush()

			if err == nil {

				_, err = r.Client.Connection.Write(
					buffer.Bytes(),
				)
			}

			if err == nil {

				r.Sent.Add(success)

				r.Batches.Add(1)
			}

			ReleaseBuffer(buffer)

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

func (r *WireBatchReplicator) Replicate(
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

func (r *WireBatchReplicator) Stop() {

	if !r.Running.Load() {
		return
	}

	r.Running.Store(false)

	close(r.Queue)

	r.WG.Wait()
}
