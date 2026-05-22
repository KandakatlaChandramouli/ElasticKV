package scheduler

import (
	"sync/atomic"
	"testing"
	"time"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/scheduler"
)

func BenchmarkFlushScheduler(
	b *testing.B,
) {

	var counter atomic.Uint64

	scheduler := engine.NewFlushScheduler(
		time.Microsecond,
		func() {
			counter.Add(1)
		},
	)

	scheduler.Start()

	time.Sleep(100 * time.Millisecond)

	scheduler.Stop()

	if counter.Load() == 0 {
		b.Fatal("scheduler failed")
	}
}
