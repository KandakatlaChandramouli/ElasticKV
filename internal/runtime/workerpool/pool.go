package workerpool

import (
	"sync"
)

type Task func()

type Pool struct {
	wg sync.WaitGroup
}

func New() *Pool {

	return &Pool{}
}

func (p *Pool) Submit(
	task Task,
) {

	p.wg.Add(1)

	go func() {

		defer p.wg.Done()

		task()
	}()
}

func (p *Pool) Wait() {

	p.wg.Wait()
}
