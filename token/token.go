package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL TERMS
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENTIFIERS
	IDENT = "IDENT"

	// LITERALS
	INT = "INT"

	// OPERATORS
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// DELMITERS
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = ")"
	RPAREN   = "("
	LBRACKET = "{"
	RBRACKET = "}"

	// KEYWORDS
	VAR      = "VAR"
	FUNCTION = "FUNCTION"
)

var keywords = map[string]TokenType{
	"var": VAR,
	"fn":  FUNCTION,
}

func LookUpIdent(literal string) TokenType {
	if tokType, exist := keywords[literal]; exist {
		return tokType
	}

	return IDENT
}
