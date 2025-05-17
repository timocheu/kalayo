package lexer

import (
	"testing"

	"github.com/timocheu/kalayo/token"
)

func TestNextToken(t *testing.T) {
	input := `;,(){}`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACKET, "{"},
		{token.RBRACKET, "}"},
		{token.EOF, ""},
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
