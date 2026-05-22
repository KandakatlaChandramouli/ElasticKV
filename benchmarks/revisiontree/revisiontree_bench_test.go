package revisiontree

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/revisiontree"
)

func BenchmarkRevisionTree(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Insert(uint64(i))
	}

	if len(runtime.Revisions) == 0 {
		b.Fatal("revision tree failed")
	}
}
