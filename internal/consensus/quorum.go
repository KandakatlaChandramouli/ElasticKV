package consensus

type Quorum struct {
	Nodes int
}

func NewQuorum(
	nodes int,
) *Quorum {

	return &Quorum{
		Nodes: nodes,
	}
}

func (q *Quorum) Majority() int {

	return (q.Nodes / 2) + 1
}

func (q *Quorum) HasQuorum(
	votes int,
) bool {

	return votes >= q.Majority()
}
