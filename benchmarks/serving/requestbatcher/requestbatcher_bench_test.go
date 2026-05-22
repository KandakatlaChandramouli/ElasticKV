package requestbatcher

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/serving/requestbatcher"
)

func BenchmarkRequestBatcher(
	b *testing.B,
) {

	requests := make([]engine.Request, 10000)

	for i := range requests {
		requests[i] = engine.Request{
			Prompt: "hello",
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = engine.Build(
			requests,
			128,
		)
	}
}
