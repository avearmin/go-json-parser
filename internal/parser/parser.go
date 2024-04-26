package parser

import (
	"errors"
	"github.com/avearmin/go-json-parser/internal/ast"
	"github.com/avearmin/go-json-parser/internal/lexer"
	"github.com/avearmin/go-json-parser/internal/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseJSON() (ast.Root, error) {
	var root ast.Root

	val, err := p.parseValue()
	if err != nil {
		return ast.Root{}, err
	}
	root.Value = &val
	return root, nil
}

func (p *Parser) parseValue() (any, error) {
	switch p.currentToken.Type {
	case token.LBrace:
		return p.parseObject()
	default:
		return p.parseLiteral()
	}
}

func (p *Parser) parseObject() (ast.Object, error) {
	object := ast.Object{Children: []ast.Property{}}

	for p.currentToken.Type != token.EOF {
		switch p.currentToken.Type {
		case token.LBrace:
			if p.peekToken.Type != token.RBrace {
				return ast.Object{}, errors.New("invalid object")
			}
		}
	}

	return object, nil
}

func (p *Parser) parseLiteral() (any, error) {
	return p.currentToken.Literal, nil
}
