package speculativeexecution

func Execute(
	primary func() string,
	speculative func() string,
) string {

	result := primary()

	if result != "" {
		return result
	}

	return speculative()
}
