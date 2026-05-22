package compactor

import (
	"github.com/KandakatlaChandramouli/ElasticKV/internal/compaction"
	"github.com/KandakatlaChandramouli/ElasticKV/internal/sstable"
)

type Runtime struct{}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Compact(
	left []sstable.Entry,
	right []sstable.Entry,
) []sstable.Entry {

	return compaction.Merge(
		left,
		right,
	)
}
