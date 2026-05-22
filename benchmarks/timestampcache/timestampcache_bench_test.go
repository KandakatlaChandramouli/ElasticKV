package timestampcache

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/timestampcache"
)

func BenchmarkTimestampCache(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Advance()
	}

	if runtime.Timestamp.Load() == 0 {
		b.Fatal("timestamp failed")
	}
}
