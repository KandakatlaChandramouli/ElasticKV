package clustersim

import (
	"math/rand"
)

type Node struct {
	ID    int
	Alive bool
}

type Cluster struct {
	Nodes []Node
}

func New(
	size int,
) *Cluster {

	cluster := &Cluster{
		Nodes: make([]Node, size),
	}

	for i := 0; i < size; i++ {

		cluster.Nodes[i] = Node{
			ID:    i,
			Alive: true,
		}
	}

	return cluster
}

func (c *Cluster) RandomFailure() {

	if len(c.Nodes) == 0 {
		return
	}

	idx := rand.Intn(
		len(c.Nodes),
	)

	c.Nodes[idx].Alive = false
}
