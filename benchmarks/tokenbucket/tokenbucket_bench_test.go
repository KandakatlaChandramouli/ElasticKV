package tokenbucket

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/tokenbucket"
)

func BenchmarkTokenBucket(
	b *testing.B,
) {

	runtime := engine.NewRuntime(
		uint64(b.N + 1000),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := runtime.Consume()

		if !ok {
			b.Fatal("consume failed")
		}
	}
}
