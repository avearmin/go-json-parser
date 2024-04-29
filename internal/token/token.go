package token

const (
	LBrace = "{"
	RBrace = "}"
	Colon  = ":"
	Comma  = ","

	String  = "STRING"
	Boolean = "BOOLEAN"
	Number  = "NUMBER"
	Null    = "NULL"

	EOF     = "EOF"
	Illegal = "ILLEGAL"
)

type Token struct {
	Type    string
	Literal string
}

func New(tokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func NewFromByte(tokenType string, char byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func NewEOF() Token {
	return Token{
		Type:    EOF,
		Literal: "",
	}
}

func LookupIdent(ident string) string {
	switch ident {
	case "true":
		return Boolean
	case "false":
		return Boolean
	case "null":
		return Null
	default:
		return Illegal
	}
}
