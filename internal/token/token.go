package token

const (
	LBrace = "{"
	RBrace = "}"

	EOF = "EOF"
)

type Token struct {
	Type    string
	Literal string
}

func NewFromByte(tokenType string, char byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(char),
	}
}
