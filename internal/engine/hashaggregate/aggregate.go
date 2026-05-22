package hashaggregate

type Aggregator struct {
	values map[string]int
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		values: make(map[string]int),
	}
}

func (a *Aggregator) Add(
	key string,
	value int,
) {
	a.values[key] += value
}

func (a *Aggregator) Result() map[string]int {
	return a.values
}
