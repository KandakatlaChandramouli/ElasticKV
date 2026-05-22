package backpressure

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/backpressure"
)

func BenchmarkBackpressure(
	b *testing.B,
) {

	runtime := engine.NewRuntime(
		uint64(b.N + 1),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := runtime.Push()

		if !ok {
			b.Fatal("push failed")
		}
	}
}
