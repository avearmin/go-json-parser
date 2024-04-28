package token

const (
	LBrace = "{"
	RBrace = "}"
	Colon  = ":"

	String = "STRING"

	EOF = "EOF"
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
