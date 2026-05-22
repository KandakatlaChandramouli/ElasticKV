package arenallocator

type Arena struct {
	buffer []byte
	offset int
}

func NewArena(
	size int,
) *Arena {

	return &Arena{
		buffer: make([]byte, size),
	}
}

func (a *Arena) Allocate(
	size int,
) []byte {

	start := a.offset
	end := start + size

	if end > len(a.buffer) {
		return nil
	}

	a.offset = end

	return a.buffer[start:end]
}
