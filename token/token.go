package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Illegal terms
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// DELMITERS
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = ")"
	RPAREN    = "("
	LBRACKET  = "{"
	RBRACKET  = "}"
)
