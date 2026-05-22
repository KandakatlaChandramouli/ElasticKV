package bufferpool

type Pool struct {
	buffers [][]byte
}

func NewPool() *Pool {
	return &Pool{
		buffers: make([][]byte, 0),
	}
}

func (p *Pool) Acquire(
	size int,
) []byte {

	return make([]byte, size)
}

func (p *Pool) Release(
	buffer []byte,
) {
	p.buffers = append(
		p.buffers,
		buffer,
	)
}
