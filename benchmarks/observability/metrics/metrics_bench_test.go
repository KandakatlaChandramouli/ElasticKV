package metrics

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/observability/metrics"
)

func BenchmarkMetrics(
	b *testing.B,
) {

	metrics := engine.NewMetrics()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		metrics.Increment()
	}
}
