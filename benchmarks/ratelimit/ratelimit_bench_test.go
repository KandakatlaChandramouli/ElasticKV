package ratelimit

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/ratelimit"
)

func BenchmarkRateLimit(
	b *testing.B,
) {

	runtime := engine.NewRuntime(
		uint64(b.N + 1),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := runtime.Allow()

		if !ok {
			b.Fatal("limit failed")
		}
	}
}
