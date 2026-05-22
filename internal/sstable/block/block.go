package block

type Entry struct {
	Key   string
	Value []byte
}

type Block struct {
	Entries []Entry
}

func New() *Block {
	return &Block{
		Entries: make([]Entry, 0),
	}
}

func (b *Block) Add(
	key string,
	value []byte,
) {

	b.Entries = append(
		b.Entries,
		Entry{
			Key:   key,
			Value: value,
		},
	)
}
