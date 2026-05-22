package memorypool

type Pool struct {
	buffers [][]float32
}

func NewPool() *Pool {
	return &Pool{
		buffers: make([][]float32, 0),
	}
}

func (p *Pool) Allocate(
	size int,
) []float32 {

	buffer := make([]float32, size)

	p.buffers = append(
		p.buffers,
		buffer,
	)

	return buffer
}

func (p *Pool) Count() int {
	return len(p.buffers)
}
