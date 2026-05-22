package checkpointer

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/checkpointer"
)

func BenchmarkCheckpointer(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Checkpoint()
	}

	if runtime.Points.Load() == 0 {
		b.Fatal("checkpoint failed")
	}
}
