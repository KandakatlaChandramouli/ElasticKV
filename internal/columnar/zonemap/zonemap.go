package zonemap

type ZoneMap struct {
	Min int
	Max int
}

func Build(
	values []int,
) ZoneMap {

	if len(values) == 0 {
		return ZoneMap{}
	}

	min := values[0]
	max := values[0]

	for _, value := range values {

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return ZoneMap{
		Min: min,
		Max: max,
	}
}
