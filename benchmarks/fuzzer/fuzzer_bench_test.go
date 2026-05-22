package fuzzer

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/fuzzer"
)

func BenchmarkRandomKeyGeneration(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	var key uint64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		key = runtime.RandomKey()
	}

	if key == 0 {
		b.Fatal("generation failed")
	}
}
