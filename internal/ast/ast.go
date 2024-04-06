package ast

type JSONVal interface {
	JSONVal()
}

type JSONInt int

func (j JSONInt) JSONVal() {}

type JSONString string

func (j JSONString) JSONVal() {}

type JSONBool bool

func (j JSONBool) JSONVal() {}

type Pair struct {
	Key   string
	Value JSONVal
}

type JSON struct {
	Pairs []Pair
}
