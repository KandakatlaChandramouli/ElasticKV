package timestamporacle

import "sync/atomic"

type Oracle struct {
	counter atomic.Uint64
}

func NewOracle() *Oracle {
	return &Oracle{}
}

func (o *Oracle) Next() uint64 {
	return o.counter.Add(1)
}
