package lexer

import "github.com/avearmin/go-json-parser/internal/token"

type Lexer struct {
	input   string
	Output  []token.Token
	pos     int
	nextPos int
	char    byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		Output: []token.Token{},
	}
	l.readChar()
	l.nextToken()

	l.lex()

	return l
}

func (l *Lexer) lex() {
	if len(l.input) == 0 {
		return
	}

	for l.Output[len(l.Output)-1].Type != token.EOF {
		l.nextToken()
	}
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

func (l *Lexer) nextToken() {
	var tok token.Token

	switch l.char {
	case '{':
		tok = token.NewFromByte(token.LBrace, l.char)
	case '}':
		tok = token.NewFromByte(token.RBrace, l.char)
	case 0:
		tok = token.NewEOF()
	}

	l.readChar()

	l.Output = append(l.Output, tok)
}
