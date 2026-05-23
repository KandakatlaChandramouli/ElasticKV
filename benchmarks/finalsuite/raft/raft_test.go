package raft

import (
	"testing"
)

type Entry struct {
	Term int
}

func BenchmarkRaftReplication(
	b *testing.B,
) {

	log := []Entry{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		log = append(
			log,
			Entry{
				Term: i,
			},
		)
	}
}
