package scattergather

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/scattergather"
)

func BenchmarkScatterGather(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Execute()
	}

	if runtime.Count == 0 {
		b.Fatal("runtime failed")
	}
}
