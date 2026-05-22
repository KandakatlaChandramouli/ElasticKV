package filter

type Bloom struct {
	keys map[string]bool
}

func New() *Bloom {
	return &Bloom{
		keys: make(map[string]bool),
	}
}

func (b *Bloom) Add(
	key string,
) {
	b.keys[key] = true
}

func (b *Bloom) MayContain(
	key string,
) bool {

	return b.keys[key]
}
