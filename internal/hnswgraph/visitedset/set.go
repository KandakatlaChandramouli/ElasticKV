package visitedset

type VisitedSet struct {
	visited map[int]struct{}
}

func NewVisitedSet() *VisitedSet {
	return &VisitedSet{
		visited: make(map[int]struct{}),
	}
}

func (v *VisitedSet) Add(id int) {
	v.visited[id] = struct{}{}
}

func (v *VisitedSet) Contains(id int) bool {
	_, ok := v.visited[id]
	return ok
}
