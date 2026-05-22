package telemetry

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/telemetry"
)

func BenchmarkTelemetry(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Record()
	}

	if runtime.Events.Load() == 0 {
		b.Fatal("telemetry failed")
	}
}
