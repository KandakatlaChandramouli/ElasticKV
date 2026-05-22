package parser

type Query struct {
	Select string
	From   string
}

func Parse(
	tokens []string,
) Query {

	query := Query{}

	if len(tokens) >= 4 {

		query.Select = tokens[1]
		query.From = tokens[3]
	}

	return query
}
