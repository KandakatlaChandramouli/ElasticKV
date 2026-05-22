package registeredbuffer

type Buffer struct {
	ID   int
	Data []byte
}

func Register(
	id int,
	size int,
) Buffer {

	return Buffer{
		ID:   id,
		Data: make([]byte, size),
	}
}
