package shutdown

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/shutdown"
)

func BenchmarkShutdown(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Stop()
	}

	if !runtime.IsStopped() {
		b.Fatal("shutdown failed")
	}
}
