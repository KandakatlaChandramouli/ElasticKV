package queryintent

func Detect(
	query string,
) string {

	if len(query) > 64 {
		return "semantic-search"
	}

	return "keyword-search"
}
