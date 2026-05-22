package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/wal"
)

func BenchmarkWALAppend(
	b *testing.B,
) {

	path := "wal/test.log"

	_ = os.Remove(path)

	segment, err := wal.OpenSegment(
		path,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer segment.Close()

	payload := bytes.Repeat(
		[]byte("W"),
		1024,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		entry := wal.Entry{
			SequenceID: uint64(i),
			Payload:    payload,
		}

		_, err := segment.Append(
			entry,
		)

		if err != nil {
			b.Fatal(err)
		}
	}

	b.StopTimer()

	err = segment.Sync()

	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkWALReplay(
	b *testing.B,
) {

	path := "wal/replay.log"

	_ = os.Remove(path)

	segment, err := wal.OpenSegment(
		path,
	)

	if err != nil {
		b.Fatal(err)
	}

	payload := bytes.Repeat(
		[]byte("R"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		entry := wal.Entry{
			SequenceID: uint64(i),
			Payload:    payload,
		}

		_, err := segment.Append(
			entry,
		)

		if err != nil {
			b.Fatal(err)
		}
	}

	err = segment.Sync()

	if err != nil {
		b.Fatal(err)
	}

	err = segment.Close()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		count := 0

		err := wal.Replay(
			path,
			func(entry wal.Entry) error {

				count++

				return nil
			},
		)

		if err != nil {
			b.Fatal(err)
		}

		if count != 100000 {
			b.Fatalf(
				"expected 100000 entries got %d",
				count,
			)
		}
	}
}
