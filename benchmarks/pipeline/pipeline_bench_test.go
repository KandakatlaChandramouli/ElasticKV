package pipeline

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/pipeline"
)

func BenchmarkPipelinePush(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Push(
			uint64(i),
		)
	}

	if len(runtime.Stages) == 0 {
		b.Fatal("pipeline failed")
	}
}
