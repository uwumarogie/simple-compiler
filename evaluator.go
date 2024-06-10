package main

import (
	"fmt"
)

type Environment struct {
	store map[string]int64
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]int64)}
}

func (e *Environment) Get(name string) (int64, bool) {
	val, ok := e.store[name]
	return val, ok
}

func (e *Environment) Set(name string, value int64) int64 {
	e.store[name] = value
	return value
}

type Evaluator struct {
	env *Environment
}

func NewEvaluator(env *Environment) *Evaluator {
	return &Evaluator{env: env}
}

func (e *Evaluator) Eval(node Node) int64 {
	switch n := node.(type) {
	case *IntegerLiteral:
		return n.Value
	case *Identifier:
		if val, ok := e.env.Get(n.Name); ok {
			return val
		} else {
			fmt.Printf("undefined variable: %s\n", n.Name)
			return 0
		}
	case *Assignment:
		val := e.Eval(n.Value)
		e.env.Set(n.Name.Name, val)
		return val
	case *InfixExpression:
		left := e.Eval(n.Left)
		right := e.Eval(n.Right)
		switch n.Operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		}
	}
	return 0
}
