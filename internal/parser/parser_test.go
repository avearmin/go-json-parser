package parser

import (
	"github.com/avearmin/go-json-parser/internal/ast"
	"github.com/avearmin/go-json-parser/internal/lexer"
	"reflect"
	"testing"
)

func TestParseJson(t *testing.T) {
	tests := map[string]struct {
		input string
		want  ast.Root
	}{
		"braces": {
			input: "{}",
			want:  ast.Root{Value: ast.Object{Children: []ast.Property{}}},
		},
	}

	for name, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		t.Run(name, func(t *testing.T) {
			got, err := p.ParseJSON()
			if err != nil {
				t.Fail()
				t.Log(err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fail()
				t.Log("ast does not match test") // gotta give this a better fail message
			}
		})
	}
}
