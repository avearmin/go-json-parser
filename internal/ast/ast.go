package ast

type NodeType int

const (
	ObjectType NodeType = iota
	ArrayType
	StringType
	BooleanType
	NullType
	ZeroValue // this is not a real node type, but indicates a failure in parsing
)

type Node interface {
	Type() NodeType
}

type Root struct {
	Value Node
}

type Property struct {
	Key   String
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

type ZeroValueNode struct{}

func (z ZeroValueNode) Type() NodeType {
	return ZeroValue
}
