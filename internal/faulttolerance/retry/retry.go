package retry

func Execute(
	attempts int,
	fn func() bool,
) bool {

	for i := 0; i < attempts; i++ {

		if fn() {
			return true
		}
	}

	return false
}
