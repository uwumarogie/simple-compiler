package main

import (
	"fmt"
)

func main() {
	input := "x = 3 + 5 * (3 - 1)"
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	ast := parser.Parse()

	if len(parser.errors) > 0 {
		for _, err := range parser.errors {
			fmt.Println(err)
		}
		return
	}

	env := NewEnvironment()
	evaluator := NewEvaluator(env)
	result := evaluator.Eval(ast)
	fmt.Printf("Result: %d\n", result)

	val, _ := env.Get("x")
	fmt.Printf("x: %d\n", val)
}
