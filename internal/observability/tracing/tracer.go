package tracing

type Span struct {
	Name string
}

type Tracer struct {
	spans []Span
}

func NewTracer() *Tracer {
	return &Tracer{
		spans: make([]Span, 0),
	}
}

func (t *Tracer) Start(
	name string,
) {

	t.spans = append(
		t.spans,
		Span{Name: name},
	)
}

func (t *Tracer) Count() int {
	return len(t.spans)
}
