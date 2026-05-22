package checksumtree

import "hash/fnv"

type Runtime struct{}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Sum(
	data []byte,
) uint64 {

	hash := fnv.New64a()

	hash.Write(data)

	return hash.Sum64()
}
