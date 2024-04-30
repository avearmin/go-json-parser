package lexer

import (
	"fmt"
	"github.com/avearmin/go-json-parser/internal/token"
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []token.Token
	}{
		"braces": {
			input: "{ }",
			want: []token.Token{
				{token.LBrace, "{"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
		"string value": {
			input: "{\"foo\":\"bar\"}",
			want: []token.Token{
				{token.LBrace, "{"},
				{token.String, "foo"},
				{token.Colon, ":"},
				{token.String, "bar"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
		"boolean value": {
			input: "{\"foo\":true}",
			want: []token.Token{
				{token.LBrace, "{"},
				{token.String, "foo"},
				{token.Colon, ":"},
				{token.Boolean, "true"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
		"number value": {
			input: "{\"foo\":6969}",
			want: []token.Token{
				{token.LBrace, "{"},
				{token.String, "foo"},
				{token.Colon, ":"},
				{token.Number, "6969"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
		"null value": {
			input: "{\"foo\":null}",
			want: []token.Token{
				{token.LBrace, "{"},
				{token.String, "foo"},
				{token.Colon, ":"},
				{token.Null, "null"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
		"array value": {
			input: "{\"foo\":[\"bar\", true, 1994, null]}",
			want: []token.Token{
				{token.LBrace, "{"},
				{token.String, "foo"},
				{token.Colon, ":"},
				{token.LBracket, "["},
				{token.String, "bar"},
				{token.Comma, ","},
				{token.Boolean, "true"},
				{token.Comma, ","},
				{token.Number, "1994"},
				{token.Comma, ","},
				{token.Null, "null"},
				{token.RBracket, "]"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
		"multiple properties": {
			input: `
				{
					"key1": true,
					"key2": false,
					"key3": null,
					"key4": "value",
					"key5": 101
				}
			`,
			want: []token.Token{
				{token.LBrace, "{"},
				{token.String, "key1"},
				{token.Colon, ":"},
				{token.Boolean, "true"},
				{token.Comma, ","},
				{token.String, "key2"},
				{token.Colon, ":"},
				{token.Boolean, "false"},
				{token.Comma, ","},
				{token.String, "key3"},
				{token.Colon, ":"},
				{token.Null, "null"},
				{token.Comma, ","},
				{token.String, "key4"},
				{token.Colon, ":"},
				{token.String, "value"},
				{token.Comma, ","},
				{token.String, "key5"},
				{token.Colon, ":"},
				{token.Number, "101"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
	}

	for name, test := range tests {
		lexer := New(test.input)
		t.Run(name, func(t *testing.T) {
			for i := range test.want {
				got := lexer.NextToken()
				fmt.Println(i)
				if !isEqualTokens(got, test.want[i]) {
					t.Fail()
					log.Printf("got=%+v, but want=%+v", got, test.want[i])
				}
			}
		})

	}
}

func isEqualTokens(tokenOne, tokenTwo token.Token) bool {
	return (tokenOne.Type == tokenTwo.Type) && (tokenOne.Literal == tokenTwo.Literal)
}
