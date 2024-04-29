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
		"simple string key/value": {
			input: "{\"foo\":\"bar\"}",
			want: ast.Root{
				Value: ast.Object{
					Children: []ast.Property{
						{"foo", ast.String{Value: "bar"}},
					},
				},
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
			want: ast.Root{
				Value: ast.Object{
					Children: []ast.Property{
						{"key1", ast.Boolean{Value: true}},
						{"key2", ast.Boolean{Value: false}},
						{"key3", ast.Null{}},
						{"key4", ast.String{Value: "value"}},
						{"key5", ast.Number{Value: 101}},
					},
				},
			},
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
				t.Logf("expected %+v, got %+v", test.want, got)
			}
		})
	}
}
