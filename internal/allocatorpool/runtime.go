package allocatorpool

import "sync"

type Runtime struct {
	Pool sync.Pool
}

func NewRuntime() *Runtime {

	return &Runtime{
		Pool: sync.Pool{
			New: func() any {
				return make(
					[]byte,
					4096,
				)
			},
		},
	}
}

func (r *Runtime) Get() []byte {

	return r.Pool.Get().([]byte)
}

func (r *Runtime) Put(
	buffer []byte,
) {

	r.Pool.Put(buffer)
}
