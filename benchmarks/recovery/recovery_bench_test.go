package recovery

import (
	"bytes"
	"os"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/recovery"
	"github.com/KandakatlaChandramouli/ElasticKV/internal/wal"
)

func BenchmarkRecoveryReplay(
	b *testing.B,
) {

	path := "recovery.log"

	_ = os.Remove(path)

	manager, err := wal.OpenManager(
		path,
		64*1024*1024,
	)

	if err != nil {
		b.Fatal(err)
	}

	payload := bytes.Repeat(
		[]byte("R"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		err := manager.Append(
			wal.Entry{
				SequenceID: uint64(i),
				Payload:    payload,
			},
		)

		if err != nil {
			b.Fatal(err)
		}
	}

	runtime := engine.NewRuntime(
		manager,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		state, err := runtime.Replay()

		if err != nil {
			b.Fatal(err)
		}

		if len(state) != 100000 {
			b.Fatal("replay failed")
		}
	}
}
