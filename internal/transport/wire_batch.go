package transport

import (
	"bufio"
	"bytes"
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

		batch := make([]ReplicationFrame, 0, r.BatchSize)

		flush := func() {

			if len(batch) == 0 {
				return
			}

			buffer := bytes.NewBuffer(
				make([]byte, 0, 1024*len(batch)),
			)

			writer := bufio.NewWriterSize(
				buffer,
				1024*1024,
			)

			success := uint64(0)

			for _, frame := range batch {

				data, err := EncodeFrame(
					frame,
				)

				if err != nil {
					continue
				}

				_, err = writer.Write(data)

				if err != nil {
					continue
				}

				success++
			}

			err := writer.Flush()

			if err != nil {
				batch = batch[:0]
				return
			}

			_, err = r.Client.Connection.Write(
				buffer.Bytes(),
			)

			if err != nil {
				batch = batch[:0]
				return
			}

			r.Sent.Add(success)

			r.Batches.Add(1)

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
