package transport

import (
	"sync"
	"sync/atomic"
)

type Replicator struct {
	Client  *Client
	Queue   chan ReplicationFrame
	Running atomic.Bool
	WG      sync.WaitGroup
	Sent    atomic.Uint64
	Dropped atomic.Uint64
}

func NewReplicator(
	client *Client,
	queueSize int,
) *Replicator {

	r := &Replicator{
		Client: client,
		Queue:  make(chan ReplicationFrame, queueSize),
	}

	r.Running.Store(true)

	return r
}

func (r *Replicator) Start() {

	r.WG.Add(1)

	go func() {

		defer r.WG.Done()

		for frame := range r.Queue {

			err := r.Client.Send(
				frame,
			)

			if err != nil {
				continue
			}

			r.Sent.Add(1)
		}
	}()
}

func (r *Replicator) Replicate(
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

func (r *Replicator) Stop() {

	if !r.Running.Load() {
		return
	}

	r.Running.Store(false)

	close(r.Queue)

	r.WG.Wait()
}
