package pipelinequeue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/pipelinequeue"
)

func BenchmarkPipelineQueue(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Push(uint64(i))
	}

	if len(runtime.Stages) == 0 {
		b.Fatal("pipeline queue failed")
	}
}
