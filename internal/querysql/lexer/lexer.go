package lexer

import "strings"

func Tokenize(
	query string,
) []string {

	return strings.Fields(
		query,
	)
}
