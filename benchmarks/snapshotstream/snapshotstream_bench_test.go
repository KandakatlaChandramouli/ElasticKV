package snapshotstream

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/snapshotstream"
)

func BenchmarkSnapshotStream(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Send(
			1024 * 1024,
		)
	}

	if runtime.Bytes.Load() == 0 {
		b.Fatal("stream failed")
	}
}
