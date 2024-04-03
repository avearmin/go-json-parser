package lexer

import (
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
		t.Run(name, func(t *testing.T) {
			for i := range test.want {
				if len(lexer.Output) != len(test.want) {
					t.Fatalf("lexer.Output=%d, but test.want=%d", len(lexer.Output), len(test.want))
				}

				if test.want[i].Type != lexer.Output[i].Type {
					t.Fatalf("lexer.Output[%d].Type=%s, but test.want[%d].Type=%s",
						i, test.want[i].Type, i, lexer.Output[i].Type)
				}

				if test.want[i].Literal != lexer.Output[i].Literal {
					t.Fatalf("lexer.Output[%d].Type=%s, but test.want[%d].Type=%s",
						i, test.want[i].Literal, i, lexer.Output[i].Literal)
				}
			}
		})

	}
}
