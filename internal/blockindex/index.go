package blockindex

type Entry struct {
	Key    uint64
	Offset uint64
}

type Index struct {
	Entries []Entry
}

func New() *Index {

	return &Index{
		Entries: make([]Entry, 0),
	}
}

func (i *Index) Add(
	key uint64,
	offset uint64,
) {

	i.Entries = append(
		i.Entries,
		Entry{
			Key:    key,
			Offset: offset,
		},
	)
}

func (i *Index) Lookup(
	key uint64,
) (uint64, bool) {

	for _, entry := range i.Entries {

		if entry.Key == key {
			return entry.Offset, true
		}
	}

	return 0, false
}
