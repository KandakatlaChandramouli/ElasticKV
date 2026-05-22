package deviceplanner

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/multigpu/deviceplanner"
)

func BenchmarkDevicePlanner(
    b *testing.B,
) {

    devices := []engine.Device{
        {ID:0, MemoryGB:16},
        {ID:1, MemoryGB:16},
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {

        _ = engine.Plan(
            devices,
            10000,
        )
    }
}
