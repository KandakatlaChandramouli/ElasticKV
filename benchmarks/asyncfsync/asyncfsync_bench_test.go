package asyncfsync

import (
	"testing"
	"time"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/asyncfsync"
)

func BenchmarkAsyncFSync(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Sync()
	}

	time.Sleep(
		100 * time.Millisecond,
	)

	if runtime.Syncs.Load() == 0 {
		b.Fatal("fsync failed")
	}
}
