package lexer

import (
	"errors"
	"github.com/avearmin/go-json-parser/internal/token"
)

type Lexer struct {
	input   string
	pos     int
	nextPos int
	char    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input + " "} // i feel this is spaghetti, but adding the whitespace to the end of input solves the
	l.NextToken()                   // case of a primitive value as the root. Without this, the last char would not
	// be read, resulting in an illegal token. On the brightside, the whitespace will
	return l // always be eaten, so it won't affect anything but that 1 case.
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.char = 0
	} else {
		l.pos = l.nextPos
		l.nextPos = l.pos + 1
		l.char = l.input[l.pos]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.consumeWhitespaces()

	switch l.char {
	case '{':
		tok = token.NewFromByte(token.LBrace, l.char)
	case '}':
		tok = token.NewFromByte(token.RBrace, l.char)
	case '[':
		tok = token.NewFromByte(token.LBracket, l.char)
	case ']':
		tok = token.NewFromByte(token.RBracket, l.char)
	case ':':
		tok = token.NewFromByte(token.Colon, l.char)
	case ',':
		tok = token.NewFromByte(token.Comma, l.char)
	case '"':
		l.readChar()
		tok = token.New(token.String, l.readString())
	case 0:
		tok = token.NewEOF()
	default:
		if isLetter(l.char) {
			ident := l.readIdent()
			identType := token.LookupIdent(ident)
			return token.New(identType, ident)
		} else if isDigit(l.char) {
			num, err := l.readNumber()
			if err != nil {
				return token.New(token.Illegal, num)
			} else {
				return token.New(token.Number, num)
			}
		} else {
			return token.NewFromByte(token.Illegal, l.char)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) consumeWhitespaces() {
	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	pos := l.pos
	for l.char != '"' {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readIdent() string {
	pos := l.pos
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() (string, error) {
	pos := l.pos
	decCount := 0

	for isDigit(l.char) {
		if l.char == '.' {
			decCount++
		}
		l.readChar()
	}

	if decCount > 1 {
		return l.input[pos:l.pos], errors.New("number has more than 1 decimal")
	}
	return l.input[pos:l.pos], nil
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z'
}

func isDigit(char byte) bool {
	return ('0' <= char && char <= '9') || char == '.'
}
