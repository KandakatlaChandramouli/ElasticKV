package search

import (
	"container/heap"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/hnswgraph/candidateheap"
)

func Search() int {

	h := candidateheap.NewHeap()

	heap.Push(h, candidateheap.Candidate{
		NodeID:   1,
		Distance: 0.1,
	})

	best := heap.Pop(h).(candidateheap.Candidate)

	return best.NodeID
}
