package chunking

func Chunk(
	tokens []string,
	size int,
) [][]string {

	chunks := make([][]string, 0)

	for i := 0; i < len(tokens); i += size {

		end := i + size

		if end > len(tokens) {
			end = len(tokens)
		}

		chunks = append(
			chunks,
			tokens[i:end],
		)
	}

	return chunks
}
