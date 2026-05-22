package memtable

type MemTable struct {
	entries map[string][]byte
}

func New() *MemTable {
	return &MemTable{
		entries: make(map[string][]byte),
	}
}

func (m *MemTable) Put(
	key string,
	value []byte,
) {
	m.entries[key] = value
}

func (m *MemTable) Get(
	key string,
) ([]byte, bool) {

	value, ok := m.entries[key]

	return value, ok
}
