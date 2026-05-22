package gpuexecutor

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/engine/gpuexecutor"
)

func BenchmarkGPUExecutor(
	b *testing.B,
) {

	executor := engine.NewExecutor(8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !executor.Execute() {
			b.Fatal("gpu execution failed")
		}
	}
}
