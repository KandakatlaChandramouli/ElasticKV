package docstore

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/docstore"
)

func BenchmarkDocStore(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ReportAllocs()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Execute()
	}

	if runtime.Count() == 0 {
		b.Fatal("runtime failure")
	}
}
