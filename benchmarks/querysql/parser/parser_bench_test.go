package parser

import (
	"testing"

	lex "github.com/KandakatlaChandramouli/ElasticKV/internal/querysql/lexer"
	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/querysql/parser"
)

func BenchmarkParser(
	b *testing.B,
) {

	query :=
		"SELECT vector FROM embeddings"

	tokens := lex.Tokenize(query)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = engine.Parse(
			tokens,
		)
	}
}
