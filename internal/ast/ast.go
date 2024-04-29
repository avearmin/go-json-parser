package ast

type NodeType int

const (
	ObjectType NodeType = iota
	ArrayType
	StringType
	BooleanType
	NumberType
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

type Number struct {
	Value float64
}

func (n Number) Type() NodeType {
	return NumberType
}

type Boolean struct {
	Value bool
}

func (b Boolean) Type() NodeType {
	return BooleanType
}

type Null struct{}

func (n Null) Type() NodeType {
	return NullType
}

type Object struct {
	Children []Property
}

func (o Object) Type() NodeType {
	return ObjectType
}
