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
		return nil, errors.New("something went wrong") // need to replace this later
	}
}

func (p *Parser) parseObject() (ast.Object, error) {
	object := ast.Object{Children: []ast.Property{}}

	p.nextToken()

	switch p.currentToken.Type {
	case token.RBrace:
		p.nextToken()
		for p.currentToken.Type != token.LBrace {
			p.parseProperty()
		}
		return object, nil
	}

	return object, nil
}

func (p *Parser) parseProperty() (ast.Property, error) {
	var property ast.Property

	if p.currentToken.Type != token.String && p.peekToken.Type != token.Colon {
		return ast.Property{}, errors.New("malformed property for Object type")
	}
	property.Key = p.currentToken.Literal

	p.nextToken()
	p.nextToken()

	value, err := p.parseValue()
	if err != nil {
		return ast.Property{}, err
	}
	property.Value = value

	return property, nil
}
