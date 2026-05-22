package sstable

func (t *Table) Lookup(
	key uint64,
) ([]byte, bool) {

	low := 0
	high := len(t.Index) - 1

	for low <= high {

		mid := (low + high) / 2

		entry := t.Index[mid]

		if entry.Key == key {

			value := t.Data[entry.Offset : entry.Offset+entry.Length]

			return value, true
		}

		if entry.Key < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return nil, false
}
