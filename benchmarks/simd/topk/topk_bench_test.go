package topk

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/simd/topk"
)

func BenchmarkTopk(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ReportAllocs()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !runtime.Execute() {
			b.Fatal("execution failed")
		}
	}

	if runtime.Count() == 0 {
		b.Fatal("invalid count")
	}
}
