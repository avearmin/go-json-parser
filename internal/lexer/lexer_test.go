package lexer

import (
	"fmt"
	"github.com/avearmin/go-json-parser/internal/token"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []token.Token
	}{
		"braces": {
			input: "{}",
			want: []token.Token{
				{token.LBrace, "{"},
				{token.RBrace, "}"},
				{token.EOF, ""},
			},
		},
	}

	for name, test := range tests {
		lexer := New(test.input)

		for i := range test.want {
			t.Run(name, func(t *testing.T) {
				if len(lexer.Output) != len(test.want) {
					fmt.Printf("lexer.Output=%d, but test.want=%d", len(lexer.Output), len(test.want))
					t.Fail()
				}

				if test.want[i].Type != lexer.Output[i].Type {
					fmt.Printf("lexer.Output[%d].Type=%d, but test.want[%d].Type=%d",
						i, test.want[i].Type, i, lexer.Output[i].Type)
					t.Fail()
				}

				if test.want[i].Literal != lexer.Output[i].Literal {
					fmt.Printf("lexer.Output[%d].Type=%d, but test.want[%d].Type=%d",
						i, test.want[i].Literal, i, lexer.Output[i].Literal)
					t.Fail()
				}
			})

		}

	}
}
