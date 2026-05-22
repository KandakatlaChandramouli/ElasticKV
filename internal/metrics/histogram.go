package metrics

import (
	"sort"
	"sync"
	"time"
)

type Histogram struct {
	mu      sync.Mutex
	samples []int64
}

func NewHistogram() *Histogram {
	return &Histogram{
		samples: make([]int64, 0, 1024),
	}
}

func (h *Histogram) Record(
	start time.Time,
) {

	latency := time.Since(start).Nanoseconds()

	h.mu.Lock()

	h.samples = append(
		h.samples,
		latency,
	)

	h.mu.Unlock()
}

func (h *Histogram) Percentile(
	p float64,
) int64 {

	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.samples) == 0 {
		return 0
	}

	data := make([]int64, len(h.samples))

	copy(data, h.samples)

	sort.Slice(
		data,
		func(i, j int) bool {
			return data[i] < data[j]
		},
	)

	index := int(
		(float64(len(data)-1) * p) / 100.0,
	)

	return data[index]
}
