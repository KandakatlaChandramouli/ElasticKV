package log

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/raft/log"
)

func BenchmarkRaftLog(
	b *testing.B,
) {

	raftlog := engine.New()

	entry := engine.Entry{
		Term:    1,
		Index:   1,
		Command: []byte("set"),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		raftlog.Append(
			entry,
		)
	}
}
