package pipeline

type Stage func() error

type Pipeline struct {
	stages []Stage
}

func New() *Pipeline {
	return &Pipeline{
		stages: make([]Stage, 0),
	}
}

func (p *Pipeline) Add(
	stage Stage,
) {
	p.stages = append(
		p.stages,
		stage,
	)
}

func (p *Pipeline) Execute() error {

	for _, stage := range p.stages {

		if err := stage(); err != nil {
			return err
		}
	}

	return nil
}
