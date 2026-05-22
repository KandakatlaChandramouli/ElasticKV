package sstable

import (
	"encoding/binary"
	"os"
	"sort"
)

type Table struct {
	Path  string
	Index []IndexEntry
	Data  []byte
}

func Build(
	path string,
	entries []Entry,
) error {

	sort.Slice(
		entries,
		func(i, j int) bool {
			return entries[i].Key <
				entries[j].Key
		},
	)

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	header := make([]byte, 16)

	binary.LittleEndian.PutUint64(
		header[0:8],
		uint64(len(entries)),
	)

	_, err = file.Write(header)

	if err != nil {
		return err
	}

	for _, entry := range entries {

		record := make(
			[]byte,
			16+len(entry.Value),
		)

		binary.LittleEndian.PutUint64(
			record[0:8],
			entry.Key,
		)

		binary.LittleEndian.PutUint64(
			record[8:16],
			uint64(len(entry.Value)),
		)

		copy(
			record[16:],
			entry.Value,
		)

		_, err = file.Write(record)

		if err != nil {
			return err
		}
	}

	return file.Sync()
}
