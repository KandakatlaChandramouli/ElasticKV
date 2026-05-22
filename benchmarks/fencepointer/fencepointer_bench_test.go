package fencepointer

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/fencepointer"
)

func BenchmarkFencePointerSearch(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	for i := 0; i < 100000; i++ {

		runtime.Add(
			uint64(i),
			uint64(i*4096),
		)
	}

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		offset, ok := runtime.Search(
			target,
		)

		if !ok {
			b.Fatal("search failed")
		}

		if offset == 0 {
			b.Fatal("invalid offset")
		}
	}
}
