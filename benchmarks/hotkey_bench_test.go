package benchmarks

import (
	"bytes"
	"sync"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/shard"
)

func BenchmarkHotKeyContention(
	b *testing.B,
) {

	manager, err := shard.NewManager(
		4,
		1024*1024*512,
		8192,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer manager.Stop()

	payload := bytes.Repeat(
		[]byte("H"),
		1024,
	)

	hotKey := uint64(42)

	b.ResetTimer()

	var wg sync.WaitGroup

	for g := 0; g < 32; g++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for i := 0; i < b.N/32; i++ {

				manager.Dispatch(
					hotKey,
					payload,
				)
			}

		}()
	}

	wg.Wait()
}
