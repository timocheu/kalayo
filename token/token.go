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

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

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
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"var":    VAR,
	"fn":     FUNCTION,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookUpIdent(literal string) TokenType {
	if tokType, exist := keywords[literal]; exist {
		return tokType
	}

	return IDENT
}
