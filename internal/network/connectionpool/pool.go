package connectionpool

type Connection struct {
	Address string
}

type Pool struct {
	connections []Connection
}

func NewPool() *Pool {
	return &Pool{
		connections: make([]Connection, 0),
	}
}

func (p *Pool) Add(
	address string,
) {

	p.connections = append(
		p.connections,
		Connection{
			Address: address,
		},
	)
}

func (p *Pool) Count() int {
	return len(p.connections)
}
