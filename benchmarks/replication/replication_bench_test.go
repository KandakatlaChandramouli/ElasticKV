package replication

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/replication"
)

func BenchmarkReplication(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Replicate(
			3,
		)
	}

	if runtime.Replicas.Load() == 0 {
		b.Fatal("replication failed")
	}
}
