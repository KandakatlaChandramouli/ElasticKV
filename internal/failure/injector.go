package failure

import (
	"math/rand"
	"time"
)

type Injector struct {
	Probability float64
	Random      *rand.Rand
}

func NewInjector(
	probability float64,
) *Injector {

	return &Injector{
		Probability: probability,
		Random: rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		),
	}
}

func (i *Injector) Fail() bool {

	if i.Probability >= 1 {
		return true
	}

	return i.Random.Float64() <
		i.Probability
}
