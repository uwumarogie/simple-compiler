package main

type Node interface{}

type IntegerLiteral struct {
	Value int64
}

type Identifier struct {
	Name string
}

type Assignment struct {
	Name  *Identifier
	Value Node
}

type InfixExpression struct {
	Left     Node
	Operator string
	Right    Node
}
