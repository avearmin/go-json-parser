package lexer

import "github.com/avearmin/go-json-parser/internal/token"

type Lexer struct {
	input   string
	pos     int
	nextPos int
	char    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextToken()

	return l
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.char = 0
		return
	}

	l.pos = l.nextPos
	l.nextPos = l.pos + 1
	l.char = l.input[l.pos]
}

func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	l.consumeWhitespaces()

	switch l.char {
	case '{':
		tok = token.NewFromByte(token.LBrace, l.char)
	case '}':
		tok = token.NewFromByte(token.RBrace, l.char)
	case 0:
		tok = token.NewEOF()
	}

	l.readChar()

	return tok
}

func (l *Lexer) consumeWhitespaces() {
	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}
