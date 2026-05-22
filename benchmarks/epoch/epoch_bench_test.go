package epoch

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/epoch"
)

func BenchmarkEpochAdvance(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Advance()
	}

	if runtime.Epoch.Load() == 0 {
		b.Fatal("epoch failed")
	}
}
