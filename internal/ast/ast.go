package ast

type NodeType int

const (
	ObjectType NodeType = iota
	ArrayType
	StringType
	BooleanType
	NullType
)

type Node interface {
	Type() NodeType
}

type Root struct {
	Value Node
}

type Property struct {
	Key   string
	Value Node
}

type String struct {
	Value string
}

func (s String) Type() NodeType {
	return StringType
}

type Object struct {
	Children []Property
}

func (o Object) Type() NodeType {
	return ObjectType
}
