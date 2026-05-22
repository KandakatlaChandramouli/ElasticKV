package repairscan

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/repairscan"
)

func BenchmarkRepairScan(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Scan()
	}

	if runtime.Scans == 0 {
		b.Fatal("scan failed")
	}
}
