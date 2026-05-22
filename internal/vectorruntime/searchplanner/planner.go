package searchplanner

type Plan struct {
	Probes int
	Beam   int
}

func BuildPlan(
	probes int,
	beam int,
) Plan {

	return Plan{
		Probes: probes,
		Beam:   beam,
	}
}
