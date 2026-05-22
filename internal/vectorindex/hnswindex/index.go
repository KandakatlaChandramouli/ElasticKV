package hnswindex

type Node struct {
	ID     int
	Vector []float32
}

type Index struct {
	nodes []Node
}

func NewIndex() *Index {
	return &Index{
		nodes: make([]Node, 0),
	}
}

func (i *Index) Insert(
	id int,
	vector []float32,
) {

	i.nodes = append(
		i.nodes,
		Node{
			ID:     id,
			Vector: vector,
		},
	)
}

func (i *Index) Count() int {
	return len(i.nodes)
}
