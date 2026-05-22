package loadbalancer

type Balancer struct {
	nodes []string
	next  int
}

func NewBalancer(
	nodes []string,
) *Balancer {

	return &Balancer{
		nodes: nodes,
	}
}

func (b *Balancer) Next() string {

	if len(b.nodes) == 0 {
		return ""
	}

	node := b.nodes[b.next]

	b.next = (b.next + 1) % len(b.nodes)

	return node
}
