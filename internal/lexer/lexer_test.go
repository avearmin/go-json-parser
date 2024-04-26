package lexer

import (
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
	}

	for name, test := range tests {
		lexer := New(test.input)
		t.Run(name, func(t *testing.T) {
			for i := range test.want {
				got := lexer.NextToken()

				if got.Type != test.want[i].Type {
					log.Printf("got.Type=%s, but want.Type=%s", got.Type, test.want[i].Type)
					t.Fail()
				}

				if got.Literal != test.want[i].Literal {
					log.Printf("got.Literal=%s, but want.Literal=%s", got.Literal, test.want[i].Literal)
					t.Fail()
				}
			}
		})

	}
}
