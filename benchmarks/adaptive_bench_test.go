package benchmarks

import (
	"sync"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/adaptive"
)

func BenchmarkAdaptiveController(
	b *testing.B,
) {

	controller := adaptive.NewController(
		10000,
	)

	var wg sync.WaitGroup

	b.ResetTimer()

	for g := 0; g < 128; g++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for i := 0; i < b.N/128; i++ {

				ok := controller.Allow()

				if ok {
					controller.Complete()
				}
			}

		}()
	}

	wg.Wait()

	b.StopTimer()

	metrics := adaptive.Snapshot(
		controller,
	)

	b.Logf(
		"Accepted = %d",
		metrics.Accepted,
	)

	b.Logf(
		"Rejected = %d",
		metrics.Rejected,
	)

	b.Logf(
		"Threshold = %d",
		metrics.Threshold,
	)
}
