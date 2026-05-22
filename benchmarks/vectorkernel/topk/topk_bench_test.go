package topk

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorkernel/topk"
)

func BenchmarkTopK(
	b *testing.B,
) {

	values := make([]float32, 10000)

	for i := 0; i < 10000; i++ {
		values[i] = float32(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		engine.Select(values, 10)
	}
}
