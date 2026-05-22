package consensus

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/consensus"
)

func BenchmarkQuorumCheck(
	b *testing.B,
) {

	quorum := engine.NewQuorum(
		5,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := quorum.HasQuorum(
			3,
		)

		if !ok {
			b.Fatal("quorum failed")
		}
	}
}
