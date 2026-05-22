package mutationlog

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/mutationlog"
)

func BenchmarkMutationLog(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Append(uint64(i))
	}

	if len(runtime.Log) == 0 {
		b.Fatal("mutation log failed")
	}
}
