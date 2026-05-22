package zerocopy

func Transfer(
	src []byte,
	dst []byte,
) int {

	copied := copy(
		dst,
		src,
	)

	return copied
}
