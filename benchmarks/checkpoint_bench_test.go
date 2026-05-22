package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/checkpoint"
)

func BenchmarkCheckpointSnapshot(
	b *testing.B,
) {

	path := "checkpoint.snapshot"

	_ = os.Remove(path)

	runtime := checkpoint.NewRuntime()

	payload := bytes.Repeat(
		[]byte("C"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		runtime.Apply(
			uint64(i),
			payload,
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		err := runtime.Snapshot(
			path,
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCheckpointRestore(
	b *testing.B,
) {

	path := "checkpoint_restore.snapshot"

	_ = os.Remove(path)

	runtime := checkpoint.NewRuntime()

	payload := bytes.Repeat(
		[]byte("R"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		runtime.Apply(
			uint64(i),
			payload,
		)
	}

	err := runtime.Snapshot(
		path,
	)

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		restore := checkpoint.NewRuntime()

		err := restore.Restore(
			path,
		)

		if err != nil {
			b.Fatal(err)
		}

		if restore.Metadata.EntryCount !=
			100000 {

			b.Fatal(
				"restore mismatch",
			)
		}
	}
}
