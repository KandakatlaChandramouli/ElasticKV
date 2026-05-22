package candidateheap

import "container/heap"

type Candidate struct {
	NodeID   int
	Distance float32
}

type CandidateHeap []Candidate

func (h CandidateHeap) Len() int {
	return len(h)
}

func (h CandidateHeap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h CandidateHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CandidateHeap) Push(x interface{}) {
	*h = append(*h, x.(Candidate))
}

func (h *CandidateHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]

	*h = old[:n-1]

	return item
}

func NewHeap() *CandidateHeap {
	h := &CandidateHeap{}
	heap.Init(h)
	return h
}
