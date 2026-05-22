package scheduler

import (
	"sync"
	"time"
)

type Task func()

type FlushScheduler struct {
	Interval time.Duration
	Task     Task
	StopChan chan struct{}
	WG       sync.WaitGroup
}

func NewFlushScheduler(
	interval time.Duration,
	task Task,
) *FlushScheduler {

	return &FlushScheduler{
		Interval: interval,
		Task:     task,
		StopChan: make(chan struct{}),
	}
}

func (f *FlushScheduler) Start() {

	f.WG.Add(1)

	go func() {

		defer f.WG.Done()

		ticker := time.NewTicker(
			f.Interval,
		)

		defer ticker.Stop()

		for {

			select {

			case <-ticker.C:
				f.Task()

			case <-f.StopChan:
				return
			}
		}
	}()
}

func (f *FlushScheduler) Stop() {

	close(f.StopChan)

	f.WG.Wait()
}
