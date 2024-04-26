package ast

type Root struct {
	Value any
}

type Property struct {
	Key   string
	Value any
}

type Object struct {
	Children []Property
}
