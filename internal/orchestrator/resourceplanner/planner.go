package resourceplanner

type Plan struct {
	CPU    int
	Memory int
	GPU    int
}

func BuildPlan(
	cpu int,
	memory int,
	gpu int,
) Plan {

	return Plan{
		CPU:    cpu,
		Memory: memory,
		GPU:    gpu,
	}
}
