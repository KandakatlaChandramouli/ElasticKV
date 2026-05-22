package objectstore

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/objectstore"
)

func BenchmarkObjectStore(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Put(uint64(i))
	}

	if len(runtime.Objects) == 0 {
		b.Fatal("object store failed")
	}
}
