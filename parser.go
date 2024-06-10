package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
	lexer  *Lexer
	tokens []Token
	pos    int
	errors []string
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) Parse() Node {
	p.tokens = []Token{}
	for tok := p.lexer.NextToken(); tok.Type != TOKEN_EOF; tok = p.lexer.NextToken() {
		p.tokens = append(p.tokens, tok)
	}

	return p.parseProgram()
}

func (p *Parser) parseProgram() Node {
	if p.match(TOKEN_IDENT) && p.match(TOKEN_ASSIGN) {
		name := &Identifier{Name: p.previous().Literal}
		value := p.parseExpression()
		return &Assignment{Name: name, Value: value}
	}
	return p.parseExpression()
}

func (p *Parser) parseExpression() Node {
	node := p.parseTerm()
	for p.match(TOKEN_PLUS, TOKEN_MINUS) {
		operator := p.previous().Literal
		right := p.parseTerm()
		node = &InfixExpression{Left: node, Operator: operator, Right: right}
	}
	return node
}

func (p *Parser) parseTerm() Node {
	node := p.parseFactor()
	for p.match(TOKEN_ASTERISK, TOKEN_SLASH) {
		operator := p.previous().Literal
		right := p.parseFactor()
		node = &InfixExpression{Left: node, Operator: operator, Right: right}
	}
	return node
}

func (p *Parser) parseFactor() Node {
	if p.match(TOKEN_INT) {
		value, _ := strconv.ParseInt(p.previous().Literal, 10, 64)
		return &IntegerLiteral{Value: value}
	} else if p.match(TOKEN_IDENT) {
		return &Identifier{Name: p.previous().Literal}
	} else if p.match(TOKEN_LPAREN) {
		node := p.parseExpression()
		if !p.match(TOKEN_RPAREN) {
			p.errors = append(p.errors, "expected closing parenthesis")
		}
		return node
	}
	p.errors = append(p.errors, fmt.Sprintf("unexpected token: %v", p.tokens[p.pos]))
	return nil
}

func (p *Parser) match(types ...string) bool {
	for _, t := range types {
		if p.check(t) {
			p.pos++
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType string) bool {
	if p.pos >= len(p.tokens) {
		return false
	}
	return p.tokens[p.pos].Type == tokenType
}

func (p *Parser) previous() Token {
	return p.tokens[p.pos-1]
}
