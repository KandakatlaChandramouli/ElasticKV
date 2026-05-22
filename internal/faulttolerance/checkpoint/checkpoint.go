package checkpoint

type Checkpoint struct {
	ID     int
	Offset uint64
}

func Create(
	id int,
	offset uint64,
) Checkpoint {

	return Checkpoint{
		ID:     id,
		Offset: offset,
	}
}
