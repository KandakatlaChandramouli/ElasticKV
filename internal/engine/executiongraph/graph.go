package executiongraph

type Node struct {
	ID       int
	Children []*Node
}

func NewNode(
	id int,
) *Node {

	return &Node{
		ID:       id,
		Children: make([]*Node, 0),
	}
}

func (n *Node) AddChild(
	child *Node,
) {
	n.Children = append(
		n.Children,
		child,
	)
}
