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
		"object root": {
			input: "{}",
			want:  ast.Root{Value: ast.Object{Children: []ast.Property{}}},
		},
		"array root": {
			input: "[\"foo\", false, null, 99999]",
			want: ast.Root{
				Value: ast.Array{
					Children: []ast.Node{
						ast.String{Value: "foo"},
						ast.Boolean{Value: false},
						ast.Null{},
						ast.Number{Value: 99999},
					},
				},
			},
		},
		"string root": {
			input: "\"root\"",
			want: ast.Root{
				Value: ast.String{Value: "root"},
			},
		},
		"number root": {
			input: "6969",
			want: ast.Root{
				Value: ast.Number{Value: 6969},
			},
		},
		"decimal number root": {
			input: "6969.6969",
			want: ast.Root{
				Value: ast.Number{Value: 6969.6969},
			},
		},
		"boolean root": {
			input: "true",
			want: ast.Root{
				Value: ast.Boolean{Value: true},
			},
		},
		"null root": {
			input: "null",
			want: ast.Root{
				Value: ast.Null{},
			},
		},
		"object root with string key/value": {
			input: "{\"foo\":\"bar\"}",
			want: ast.Root{
				Value: ast.Object{
					Children: []ast.Property{
						{"foo", ast.String{Value: "bar"}},
					},
				},
			},
		},
		"object root with multiple properties": {
			input: `
				{
					"key1": true,
					"key2": false,
					"key3": null,
					"key4": "value",
					"key5": 101,
					"key6": 101.6969
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
						{"key6", ast.Number{Value: 101.6969}},
					},
				},
			},
		},
		"object root with array value": {
			input: "{\"foo\":[\"bar\", true, 1994, null]}",
			want: ast.Root{
				Value: ast.Object{
					Children: []ast.Property{
						{"foo", ast.Array{
							Children: []ast.Node{
								ast.String{Value: "bar"},
								ast.Boolean{Value: true},
								ast.Number{Value: 1994},
								ast.Null{},
							},
						}},
					},
				},
			},
		},
		"object root with nested objects": {
			input: `
				{
					"key1": {"nested1": {"nested2": {}}},
					"key2": {"nested3": {}}
				}
			`,
			want: ast.Root{
				Value: ast.Object{
					Children: []ast.Property{
						{"key1", ast.Object{
							Children: []ast.Property{
								{"nested1", ast.Object{
									Children: []ast.Property{
										{"nested2", ast.Object{
											Children: []ast.Property{},
										}},
									},
								}},
							},
						}},
						{"key2", ast.Object{
							Children: []ast.Property{
								{"nested3", ast.Object{
									Children: []ast.Property{},
								}},
							},
						}},
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
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("expected %+v, got %+v", test.want, got)
			}
		})
	}
}
