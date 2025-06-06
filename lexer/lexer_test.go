package lexer

import (
	"testing"

	"github.com/timocheu/kalayo/token"
)

func TestNextToken(t *testing.T) {
	input := `
	var age = 10;

	fn check(age) {
	age + 10;
	};

	var result = check(age);

	!*/<>;

	<=
	>=
	!=
	==
	`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "age"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "check"},
		{token.LPAREN, "("},
		{token.IDENT, "age"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.IDENT, "age"},
		{token.PLUS, "+"},
		{token.INT, "10"},
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
		{token.BANG, "!"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.SEMICOLON, ";"},
		{token.LT_EQUAL, "<="},
		{token.GT_EQUAL, ">="},
		{token.NOT_EQUAL, "!="},
		{token.EQUAL, "=="},
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
