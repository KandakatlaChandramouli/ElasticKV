package groupcommit

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/groupcommit"
)

func BenchmarkGroupCommit(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Commit(
			64,
		)
	}

	if runtime.Commits.Load() == 0 {
		b.Fatal("commit failed")
	}
}
