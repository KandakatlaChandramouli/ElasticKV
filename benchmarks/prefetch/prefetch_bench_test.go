package prefetch

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/prefetch"
)

func BenchmarkPrefetch(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Fetch()
	}

	if runtime.Prefetches.Load() == 0 {
		b.Fatal("prefetch failed")
	}
}
