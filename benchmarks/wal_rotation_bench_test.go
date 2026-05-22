package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/wal"
)

func BenchmarkWALRotation(
	b *testing.B,
) {

	path := "wal_rotation"

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
		[]byte("R"),
		1024,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

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

	b.StopTimer()

	err = manager.Sync()

	if err != nil {
		b.Fatal(err)
	}

	segments, err := manager.Segments()

	if err != nil {
		b.Fatal(err)
	}

	b.Logf(
		"Segments Created = %d",
		len(segments),
	)
}

func BenchmarkWALMultiReplay(
	b *testing.B,
) {

	path := "wal_replay"

	_ = os.RemoveAll(path)

	manager, err := wal.OpenManager(
		path,
		2*1024*1024,
	)

	if err != nil {
		b.Fatal(err)
	}

	payload := bytes.Repeat(
		[]byte("M"),
		1024,
	)

	for i := 0; i < 200000; i++ {

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

	err = manager.Close()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		count := 0

		err := manager.Replay(
			func(entry wal.Entry) error {

				count++

				return nil
			},
		)

		if err != nil {
			b.Fatal(err)
		}

		if count != 200000 {

			b.Fatalf(
				"expected 200000 entries got %d",
				count,
			)
		}
	}
}
