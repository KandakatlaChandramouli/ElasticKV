package repair

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/repair"
)

func BenchmarkRepairRuntime(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Repair()
	}

	if runtime.Repairs.Load() == 0 {
		b.Fatal("repair failed")
	}
}
