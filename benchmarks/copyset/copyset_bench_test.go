package copyset

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/copyset"
)

func BenchmarkCopySet(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Add(uint64(i % 5))
	}

	if len(runtime.Replicas) == 0 {
		b.Fatal("copyset failed")
	}
}
