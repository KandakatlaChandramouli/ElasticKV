package adaptive

import (
	"time"
)

type RuntimeMetrics struct {
	Load       uint64
	Accepted   uint64
	Rejected   uint64
	Threshold  uint64
	RecordedAt time.Time
}

func Snapshot(
	c *Controller,
) RuntimeMetrics {

	return RuntimeMetrics{
		Load:       c.CurrentLoad.Load(),
		Accepted:   c.Accepted.Load(),
		Rejected:   c.Rejected.Load(),
		Threshold:  c.RejectThreshold.Load(),
		RecordedAt: time.Now(),
	}
}
