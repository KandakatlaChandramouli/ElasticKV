package hnsw

import (
	"math/rand"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/hnsw/core"
)

func BenchmarkHNSWRecall(
	b *testing.B,
) {

	index := engine.New()

	for i := 0; i < 100000; i++ {

		vec := make([]float32, 768)

		for j := range vec {
			vec[j] = rand.Float32()
		}

		index.Insert(i, vec)
	}

	query := make([]float32, 768)

	for i := range query {
		query[i] = rand.Float32()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		index.Search(
			query,
			10,
		)
	}
}
