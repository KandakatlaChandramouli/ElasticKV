package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/wal"
)

func BenchmarkWALIndexing(
	b *testing.B,
) {

	path := "wal_index"

	_ = os.RemoveAll(path)

	manager, err := wal.OpenManager(
		path,
		4*1024*1024,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer manager.Close()

	payload := bytes.Repeat(
		[]byte("I"),
		1024,
	)

	for i := 0; i < 1000000; i++ {

		entry := wal.Entry{
			SequenceID: uint64(i),
			Payload:    payload,
		}

		err := manager.Append(
			entry,
		)

		if err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		target := uint64(
			i % 1000000,
		)

		entry, ok := manager.Find(
			target,
		)

		if !ok {
			b.Fatal("index miss")
		}

		_ = entry
	}

	b.StopTimer()

	b.Logf(
		"Indexed Entries = %d",
		len(manager.Index.Entries),
	)
}
