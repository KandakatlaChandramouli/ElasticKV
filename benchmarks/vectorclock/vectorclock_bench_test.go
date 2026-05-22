package vectorclock

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorclock"
)

func BenchmarkVectorClock(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Tick(
			uint64(i % 16),
		)
	}

	if len(runtime.Clocks) == 0 {
		b.Fatal("clock failed")
	}
}
