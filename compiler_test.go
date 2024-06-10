package main

import (
	"testing"
)

func TestLexer(t *testing.T) {
	input := "x = 3 + 5 * (2 - 1)"
	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{TOKEN_IDENT, "x"},
		{TOKEN_ASSIGN, "="},
		{TOKEN_INT, "3"},
		{TOKEN_PLUS, "+"},
		{TOKEN_INT, "5"},
		{TOKEN_ASTERISK, "*"},
		{TOKEN_LPAREN, "("},
		{TOKEN_INT, "2"},
		{TOKEN_MINUS, "-"},
		{TOKEN_INT, "1"},
		{TOKEN_RPAREN, ")"},
		{TOKEN_EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
