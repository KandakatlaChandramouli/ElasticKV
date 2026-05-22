package checkpointdelta

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/checkpointdelta"
)

func BenchmarkCheckpointDelta(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.ApplyDelta(
			64,
		)
	}

	if runtime.Deltas.Load() == 0 {
		b.Fatal("delta failed")
	}
}
