package dotproduct

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorkernel/dotproduct"
)

func BenchmarkDotProduct(
	b *testing.B,
) {

	a := make([]float32, 1024)
	c := make([]float32, 1024)

	for i := 0; i < 1024; i++ {
		a[i] = float32(i)
		c[i] = float32(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		engine.Compute(a, c)
	}
}
