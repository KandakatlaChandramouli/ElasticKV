package metrics

import "sync/atomic"

type Metrics struct {
	requests atomic.Uint64
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) Increment() {
	m.requests.Add(1)
}

func (m *Metrics) Count() uint64 {
	return m.requests.Load()
}
