package hnswindex

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorindex/hnswindex"
)

func BenchmarkHNSWIndex(
	b *testing.B,
) {

	index := engine.NewIndex()

	vector := make([]float32, 768)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		index.Insert(
			i,
			vector,
		)
	}
}
