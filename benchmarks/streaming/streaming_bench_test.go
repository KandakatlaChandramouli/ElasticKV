package streaming

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/streaming"
)

func BenchmarkStreaming(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Stream(
			4096,
		)
	}

	if runtime.Bytes.Load() == 0 {
		b.Fatal("stream failed")
	}
}
