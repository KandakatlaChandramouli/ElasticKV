package network

func CopyInto(
	dst []byte,
	src []byte,
) {

	copy(
		dst,
		src,
	)
}

func SliceView(
	buffer []byte,
	offset uint64,
	length uint64,
) []byte {

	return buffer[offset : offset+length]
}
