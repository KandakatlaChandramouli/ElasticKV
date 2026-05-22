package batchdistance

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/gpu/batchdistance"
)

func BenchmarkBatchDistance(
	b *testing.B,
) {

	vectors := make([][]float32, 1024)

	for i := range vectors {
		vectors[i] = make([]float32, 768)
	}

	query := make([]float32, 768)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = engine.Compute(
			vectors,
			query,
		)
	}
}
