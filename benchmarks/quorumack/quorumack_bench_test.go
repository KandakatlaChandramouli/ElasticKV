package quorumack

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/quorumack"
)

func BenchmarkQuorumAck(
	b *testing.B,
) {

	runtime := engine.NewRuntime(
		3,
	)

	runtime.Ack()
	runtime.Ack()
	runtime.Ack()

	if !runtime.Quorum() {
		b.Fatal("quorum failed")
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Ack()
	}
}
