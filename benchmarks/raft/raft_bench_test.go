package raft

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/raft"
)

func BenchmarkRaftAppend(
	b *testing.B,
) {

	log := engine.NewLog()

	payload := bytes.Repeat(
		[]byte("R"),
		1024,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		log.Append(
			engine.Entry{
				Index: uint64(i),
				Term:  1,
				Data:  payload,
			},
		)
	}

	if log.LastIndex() == 0 {
		b.Fatal("append failed")
	}
}
