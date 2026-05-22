package benchmarks

import (
	"bytes"
	"sync"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/shard"
)

func BenchmarkBackpressure(
	b *testing.B,
) {

	manager, err := shard.NewManager(
		4,
		1024*1024*256,
		64,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer manager.Stop()

	payload := bytes.Repeat(
		[]byte("B"),
		1024,
	)

	var wg sync.WaitGroup

	b.ResetTimer()

	for g := 0; g < 128; g++ {

		wg.Add(1)

		go func(offset int) {

			defer wg.Done()

			for i := 0; i < b.N/128; i++ {

				manager.Dispatch(
					uint64(i+offset),
					payload,
				)
			}

		}(g * 100000)
	}

	wg.Wait()

	b.StopTimer()

	b.Logf(
		"Dropped Requests = %d",
		manager.Dropped.Load(),
	)
}
