package snapshotdelta

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/snapshotdelta"
)

func BenchmarkSnapshotDelta(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Append(
			uint64(i),
		)
	}

	if len(runtime.Deltas) == 0 {
		b.Fatal("delta failed")
	}
}
