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
	root.Value = val
	return root, nil
}

func (p *Parser) parseValue() (ast.Node, error) {
	switch p.currentToken.Type {
	case token.LBrace:
		return p.parseObject()
	default:
		return ast.ZeroValueNode{}, errors.New("something went wrong") // need to replace this later
	}
}

func (p *Parser) parseObject() (ast.Object, error) {
	object := ast.Object{Children: []ast.Property{}}

	p.nextToken()

	switch p.currentToken.Type {
	case token.RBrace:
		return object, nil
	}

	return object, nil
}
