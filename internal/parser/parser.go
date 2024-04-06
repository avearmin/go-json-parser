package parser

import (
	"github.com/avearmin/go-json-parser/internal/lexer"
	"github.com/avearmin/go-json-parser/internal/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}
