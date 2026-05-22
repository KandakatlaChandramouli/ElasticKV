package gpuexecutor

type GPUExecutor struct {
	Streams int
}

func NewExecutor(
	streams int,
) *GPUExecutor {

	return &GPUExecutor{
		Streams: streams,
	}
}

func (g *GPUExecutor) Execute() bool {
	return g.Streams > 0
}
