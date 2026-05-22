package core

import (
	"math"
	"sort"
)

type Neighbor struct {
	ID    int
	Score float32
}

type Node struct {
	ID        int
	Vector    []float32
	Neighbors []int
}

type HNSW struct {
	nodes map[int]*Node
}

func New() *HNSW {
	return &HNSW{
		nodes: make(map[int]*Node),
	}
}

func cosine(
	a []float32,
	b []float32,
) float32 {

	var dot float32
	var normA float32
	var normB float32

	for i := range a {

		dot += a[i] * b[i]

		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	return dot /
		(float32(math.Sqrt(float64(normA)))*
			float32(math.Sqrt(float64(normB))) +
			1e-9)
}

func (h *HNSW) Insert(
	id int,
	vector []float32,
) {

	node := &Node{
		ID:        id,
		Vector:    vector,
		Neighbors: make([]int, 0),
	}

	for existingID, existing := range h.nodes {

		score := cosine(
			vector,
			existing.Vector,
		)

		if score > 0.8 {

			node.Neighbors = append(
				node.Neighbors,
				existingID,
			)

			existing.Neighbors = append(
				existing.Neighbors,
				id,
			)
		}
	}

	h.nodes[id] = node
}

func (h *HNSW) Search(
	query []float32,
	k int,
) []Neighbor {

	results := make([]Neighbor, 0)

	for id, node := range h.nodes {

		score := cosine(
			query,
			node.Vector,
		)

		results = append(
			results,
			Neighbor{
				ID:    id,
				Score: score,
			},
		)
	}

	sort.Slice(
		results,
		func(i, j int) bool {
			return results[i].Score >
				results[j].Score
		},
	)

	if k > len(results) {
		k = len(results)
	}

	return results[:k]
}
