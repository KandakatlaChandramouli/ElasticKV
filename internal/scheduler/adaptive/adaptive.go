package adaptive

func Adjust(
	load int,
) int {

	if load > 1000 {
		return 32
	}

	if load > 100 {
		return 16
	}

	return 4
}
