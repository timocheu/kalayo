package lexer

import (
	"testing"

	"github.com/timocheu/kalayo/token"
)

func TestNextToken(t *testing.T) {
	input := `
	var age = ;

	fn check(age) {
	age + ;
	};

	var result = check(age);
	`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "age"},
		{token.ASSIGN, "="},
		// {token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "check"},
		{token.LPAREN, "("},
		{token.IDENT, "age"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.IDENT, "age"},
		{token.ADD, "+"},
		// {token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.RBRACKET, "}"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "check"},
		{token.LPAREN, "("},
		{token.IDENT, "age"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range test {
		currTok := l.NextToken()

		if currTok.Type != tt.expectedType {
			t.Fatalf("test[%d] Token type wrong, want: %q, got: %q",
				i, tt.expectedType, currTok.Type)
		}

		if currTok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] Token literal wrong, want: %q, got: %q",
				i, tt.expectedLiteral, currTok.Literal)
		}
	}
}
