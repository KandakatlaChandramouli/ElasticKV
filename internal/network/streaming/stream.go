package streaming

type Stream struct {
	chunks [][]byte
}

func NewStream() *Stream {
	return &Stream{
		chunks: make([][]byte, 0),
	}
}

func (s *Stream) Send(
	chunk []byte,
) {
	s.chunks = append(
		s.chunks,
		chunk,
	)
}

func (s *Stream) Count() int {
	return len(s.chunks)
}
