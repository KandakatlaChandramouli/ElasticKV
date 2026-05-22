package truncation

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/truncation"
)

func BenchmarkLogTruncation(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	for i := 0; i < 100000; i++ {

		runtime.Append(
			uint64(i),
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Truncate(
			50000,
		)
	}
}
