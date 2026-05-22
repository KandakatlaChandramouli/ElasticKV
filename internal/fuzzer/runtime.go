package fuzzer

import (
	"math/rand"
	"time"
)

type Runtime struct {
	Random *rand.Rand
}

func NewRuntime() *Runtime {

	return &Runtime{
		Random: rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		),
	}
}

func (r *Runtime) RandomKey() uint64 {

	return r.Random.Uint64()
}
