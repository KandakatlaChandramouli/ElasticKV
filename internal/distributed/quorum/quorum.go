package quorum

func HasQuorum(
	acknowledgements int,
	replicas int,
) bool {

	return acknowledgements >= (replicas/2)+1
}
