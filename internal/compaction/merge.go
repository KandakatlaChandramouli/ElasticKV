package compaction

import (
	"github.com/KandakatlaChandramouli/ElasticKV/internal/sstable"
)

func Merge(
	a []sstable.Entry,
	b []sstable.Entry,
) []sstable.Entry {

	result := make(
		[]sstable.Entry,
		0,
		len(a)+len(b),
	)

	i := 0
	j := 0

	for i < len(a) && j < len(b) {

		if a[i].Key < b[j].Key {

			result = append(result, a[i])

			i++

			continue
		}

		if a[i].Key > b[j].Key {

			result = append(result, b[j])

			j++

			continue
		}

		result = append(result, b[j])

		i++
		j++
	}

	for i < len(a) {

		result = append(result, a[i])

		i++
	}

	for j < len(b) {

		result = append(result, b[j])

		j++
	}

	return result
}
