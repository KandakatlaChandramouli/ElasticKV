package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/wal"
)

func BenchmarkWALIndexedReplay(
	b *testing.B,
) {

	path := "wal_seek"

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
		[]byte("S"),
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

	err = manager.Sync()

	if err != nil {
		b.Fatal(err)
	}

	target := uint64(900000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		count := 0

		err := manager.ReplayFrom(
			target,
			func(entry wal.Entry) error {

				count++

				return nil
			},
		)

		if err != nil {
			b.Fatal(err)
		}

		if count == 0 {

			b.Fatal(
				"expected replay entries",
			)
		}
	}

	b.StopTimer()

	indexEntry, ok := manager.Find(
		target,
	)

	if !ok {
		b.Fatal("missing sparse index")
	}

	b.Logf(
		"Seek Segment = %d",
		indexEntry.SegmentID,
	)

	b.Logf(
		"Seek Offset = %d",
		indexEntry.Offset,
	)
}
