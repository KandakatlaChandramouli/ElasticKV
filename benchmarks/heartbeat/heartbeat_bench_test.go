package heartbeat

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/heartbeat"
)

func BenchmarkHeartbeat(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Beat()
	}

	if runtime.Count() == 0 {
		b.Fatal("heartbeat failed")
	}
}
