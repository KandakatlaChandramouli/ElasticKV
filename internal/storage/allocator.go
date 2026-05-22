package storage

import "sync/atomic"

type Allocator struct {
	next uint64
}

func NewAllocator() *Allocator {
	return &Allocator{}
}

func (a *Allocator) Allocate() uint64 {
	return atomic.AddUint64(&a.next, BlockSize) - BlockSize
}
