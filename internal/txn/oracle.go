package txn

import "sync/atomic"

type Oracle struct {
	Current atomic.Uint64
}

func NewOracle() *Oracle {

	return &Oracle{}
}

func (o *Oracle) Next() uint64 {

	return o.Current.Add(1)
}

func (o *Oracle) Read() uint64 {

	return o.Current.Load()
}
