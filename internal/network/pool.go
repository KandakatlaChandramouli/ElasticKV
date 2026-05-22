package network

import "sync"

type BufferPool struct {
	Pool sync.Pool
}

func NewPool(
	size int,
) *BufferPool {

	return &BufferPool{
		Pool: sync.Pool{
			New: func() any {
				return make([]byte, size)
			},
		},
	}
}

func (b *BufferPool) Get() []byte {
	return b.Pool.Get().([]byte)
}

func (b *BufferPool) Put(
	buffer []byte,
) {
	b.Pool.Put(buffer)
}
