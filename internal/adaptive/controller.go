package adaptive

import (
	"sync/atomic"
)

type Controller struct {
	RejectThreshold atomic.Uint64
	CurrentLoad     atomic.Uint64
	Rejected        atomic.Uint64
	Accepted        atomic.Uint64
}

func NewController(
	initialThreshold uint64,
) *Controller {

	c := &Controller{}

	c.RejectThreshold.Store(
		initialThreshold,
	)

	return c
}

func (c *Controller) Allow() bool {

	load := c.CurrentLoad.Load()

	threshold := c.RejectThreshold.Load()

	if load >= threshold {

		c.Rejected.Add(1)

		return false
	}

	c.Accepted.Add(1)

	c.CurrentLoad.Add(1)

	return true
}

func (c *Controller) Complete() {

	current := c.CurrentLoad.Load()

	if current > 0 {
		c.CurrentLoad.Add(^uint64(0))
	}
}

func (c *Controller) AdjustThreshold(
	newThreshold uint64,
) {
	c.RejectThreshold.Store(
		newThreshold,
	)
}
