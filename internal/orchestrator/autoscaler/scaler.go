package autoscaler

func Scale(
	replicas int,
	load float64,
) int {

	if load > 0.80 {
		return replicas + 2
	}

	if load < 0.20 && replicas > 1 {
		return replicas - 1
	}

	return replicas
}
