package physicalplan

type Operator struct {
	Name string
}

type Plan struct {
	Operators []Operator
}

func NewPlan() *Plan {
	return &Plan{
		Operators: make([]Operator, 0),
	}
}

func (p *Plan) Add(
	name string,
) {
	p.Operators = append(
		p.Operators,
		Operator{Name: name},
	)
}
