package eviction

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/eviction"
)

func BenchmarkEviction(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	for i := 0; i < 100000; i++ {

		runtime.Touch(
			uint64(i),
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := runtime.Evict()

		if !ok {
			b.Fatal("eviction failed")
		}
	}
}
