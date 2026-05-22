package tokenstream

type Stream struct {
	tokens []string
}

func NewStream() *Stream {
	return &Stream{
		tokens: make([]string, 0),
	}
}

func (s *Stream) Push(
	token string,
) {
	s.tokens = append(
		s.tokens,
		token,
	)
}

func (s *Stream) Count() int {
	return len(s.tokens)
}
