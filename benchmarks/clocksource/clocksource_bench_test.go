package clocksource

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/clocksource"
)

func BenchmarkClockSource(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Tick()
	}

	if runtime.Ticks.Load() == 0 {
		b.Fatal("clock failed")
	}
}
