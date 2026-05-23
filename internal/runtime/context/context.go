package context

type QueryContext struct {
	Query      string
	Vector     []float32
	TopK       int
	EnableGPU  bool
	Distributed bool
}
