package parser

import (
	"errors"
	"fmt"
	"github.com/avearmin/go-json-parser/internal/ast"
	"github.com/avearmin/go-json-parser/internal/lexer"
	"github.com/avearmin/go-json-parser/internal/token"
	"strconv"
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
		p.nextToken()
		return p.parseObject()
	case token.LBracket:
		p.nextToken()
		return p.parseArray()
	case token.String:
		return ast.String{Value: p.currentToken.Literal}, nil
	case token.Number:
		return p.parseNumber()
	case token.Boolean:
		return p.parseBoolean()
	case token.Null:
		return ast.Null{}, nil
	default:
		return nil, errors.New("something went wrong") // need to replace this later
	}
}

func (p *Parser) parseObject() (ast.Object, error) {
	object := ast.Object{Children: []ast.Property{}}

	for p.currentToken.Type != token.RBrace {
		switch p.currentToken.Type {
		case token.EOF:
			return ast.Object{}, errors.New("unexpected 'EOF'")
		case token.Comma:
			if p.peekToken.Type != token.String {
				errMsg := fmt.Sprintf("unexpected '%s' when expecting 'STRING'", p.peekToken.Type)
				return ast.Object{}, errors.New(errMsg)
			}
			p.nextToken()
			continue
		}

		property, err := p.parseProperty()
		if err != nil {
			return ast.Object{}, err
		}

		object.Children = append(object.Children, property)
	}

	return object, nil
}

func (p *Parser) parseArray() (ast.Array, error) {
	array := ast.Array{Children: []ast.Node{}}
	fmt.Println(p.currentToken)
	for p.currentToken.Type != token.RBracket {
		switch p.currentToken.Type {
		case token.EOF:
			return ast.Array{}, errors.New("unexpected 'EOF'")
		case token.Comma:
			if !p.peekToken.IsValueType() && p.peekToken.Type != token.LBrace && p.peekToken.Type != token.LBracket {
				errMsg := fmt.Sprintf("unexpected '%s' when expecting Node", p.peekToken.Type)
				return ast.Array{}, errors.New(errMsg)
			}
			p.nextToken()
			continue
		}

		value, err := p.parseValue()
		if err != nil {
			return ast.Array{}, err
		}

		array.Children = append(array.Children, value)
		p.nextToken()
	}

	return array, nil
}

func (p *Parser) parseNumber() (ast.Number, error) {
	num, err := strconv.ParseFloat(p.currentToken.Literal, 64)
	if err != nil {
		return ast.Number{}, errors.New("couldn't parse number: " + p.currentToken.Literal)
	}
	return ast.Number{Value: num}, nil
}

func (p *Parser) parseBoolean() (ast.Boolean, error) {
	boolean, err := strconv.ParseBool(p.currentToken.Literal)
	if err != nil {
		return ast.Boolean{}, errors.New("couldn't parse boolean: " + p.currentToken.Literal)
	}
	return ast.Boolean{Value: boolean}, nil
}

func (p *Parser) parseProperty() (ast.Property, error) {
	var property ast.Property

	if p.currentToken.Type != token.String || p.peekToken.Type != token.Colon {
		return ast.Property{}, errors.New("malformed property for Object type")
	}
	property.Key = p.currentToken.Literal

	p.nextToken() // advance off the string
	p.nextToken() // advance off the colon

	value, err := p.parseValue()
	if err != nil {
		return ast.Property{}, err
	}
	property.Value = value

	p.nextToken() // advance off the value

	return property, nil
}
