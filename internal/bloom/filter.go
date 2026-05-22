package bloom

type Filter struct {
	Bits []bool
	Size uint64
}

func New(
	size uint64,
) *Filter {

	return &Filter{
		Bits: make([]bool, size),
		Size: size,
	}
}

func hash(
	key uint64,
	size uint64,
) uint64 {

	return (key * 11400714819323198485) % size
}

func (f *Filter) Add(
	key uint64,
) {

	index := hash(
		key,
		f.Size,
	)

	f.Bits[index] = true
}

func (f *Filter) MayContain(
	key uint64,
) bool {

	index := hash(
		key,
		f.Size,
	)

	return f.Bits[index]
}
