package remotememory

type Region struct {
	Address uint64
	Length  uint64
}

func Map(
	address uint64,
	length uint64,
) Region {

	return Region{
		Address: address,
		Length:  length,
	}
}
