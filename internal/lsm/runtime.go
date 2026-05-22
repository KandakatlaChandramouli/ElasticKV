package lsm

import (
	"fmt"
	"sync/atomic"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/sstable"
)

type Runtime struct {
	Active  *MemTable
	Flushes atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{
		Active: NewMemTable(),
	}
}

func (r *Runtime) Put(
	key uint64,
	value []byte,
) {

	r.Active.Put(
		key,
		value,
	)
}

func (r *Runtime) Flush() error {

	id := r.Flushes.Add(1)

	path := fmt.Sprintf(
		"segment_%d.sst",
		id,
	)

	entries := r.Active.Freeze()

	err := sstable.Build(
		path,
		entries,
	)

	if err != nil {
		return err
	}

	r.Active = NewMemTable()

	return nil
}
