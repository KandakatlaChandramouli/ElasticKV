package benchmarks

import (
	"bytes"
	"sync"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/shard"
)

func BenchmarkShardContention(b *testing.B) {

	manager, err := shard.NewManager(
		4,
		1024*1024*256,
		4096,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer manager.Stop()

	payload := bytes.Repeat([]byte("X"), 1024)

	b.ResetTimer()

	var wg sync.WaitGroup

	for g := 0; g < 16; g++ {

		wg.Add(1)

		go func(offset int) {

			defer wg.Done()

			for i := 0; i < b.N/16; i++ {

				manager.Dispatch(
					uint64(i+offset),
					payload,
				)
			}

		}(g * 1000000)
	}

	wg.Wait()
}
