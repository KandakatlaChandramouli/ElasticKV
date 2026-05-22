package snapshotter

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/snapshotter"
)

func BenchmarkSnapshotRuntime(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Snapshot()
	}

	if runtime.Snapshots.Load() == 0 {
		b.Fatal("snapshot failed")
	}
}
