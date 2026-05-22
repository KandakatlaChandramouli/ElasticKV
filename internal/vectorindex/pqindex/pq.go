package pqindex

type PQ struct {
	Codes [][]uint8
}

func NewPQ() *PQ {
	return &PQ{
		Codes: make([][]uint8, 0),
	}
}

func (p *PQ) Add(
	code []uint8,
) {
	p.Codes = append(
		p.Codes,
		code,
	)
}

func (p *PQ) Count() int {
	return len(p.Codes)
}
