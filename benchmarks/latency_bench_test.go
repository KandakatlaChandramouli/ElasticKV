package benchmarks

import (
	"bytes"
	"testing"
	"time"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/metrics"
	"github.com/KandakatlaChandramouli/ElasticKV/internal/storage"
)

func BenchmarkLatencyDistribution(
	b *testing.B,
) {

	engineSize := 1024 * 1024 * 1024

	engine, err := storage.Open(
		"latency.db",
		engineSize,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer engine.Close()

	payload := bytes.Repeat(
		[]byte("L"),
		1024,
	)

	maxWrites := engineSize / storage.BlockSize

	histogram := metrics.NewHistogram()

	b.ResetTimer()

	for i := 0; i < b.N && i < maxWrites; i++ {

		start := time.Now()

		_, err := engine.Write(
			uint64(i),
			payload,
		)

		if err != nil {
			b.Fatal(err)
		}

		histogram.Record(start)
	}

	b.StopTimer()

	b.Logf(
		"P50 = %d ns",
		histogram.Percentile(50),
	)

	b.Logf(
		"P95 = %d ns",
		histogram.Percentile(95),
	)

	b.Logf(
		"P99 = %d ns",
		histogram.Percentile(99),
	)

	b.Logf(
		"P99.9 = %d ns",
		histogram.Percentile(99.9),
	)
}
