package barrier

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/barrier"
)

func BenchmarkBarrier(
	b *testing.B,
) {

	runtime := engine.NewRuntime(
		1,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := runtime.Arrive()

		if !ok {
			b.Fatal("barrier failed")
		}
	}
}
