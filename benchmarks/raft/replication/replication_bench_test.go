package replication

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/raft/replication"
)

func BenchmarkReplication(
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
