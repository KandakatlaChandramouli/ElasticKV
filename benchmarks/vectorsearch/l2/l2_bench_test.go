package l2

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorsearch/l2"
)

func BenchmarkL2(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !runtime.Execute() {
			b.Fatal("execution failed")
		}
	}
}
